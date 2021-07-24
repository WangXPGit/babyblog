<template>
    <div>
        <h3>分类列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addCateVisible = true">新增分类</a-button>
                </a-col>
            </a-row>

            <a-table 
                rowKey="id" :columns="columns" :pagination='pagination' :dataSource="Catelist" bordered @change="handleTableChange">
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" icon="edit" style="margin-right:15px" @click="editCate(data.ID)">编辑</a-button>
                        <a-button type="danger" icon="delete" @click="deleteCate(data.ID)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
       
        <!-- 新增分类区域 -->
        <a-modal title="新增分类"
                :visible="addCateVisible"
                @ok="addCateOk"
                @cancel="addCateCancel"
                :destroyOnClose="true">
            <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
                <a-form-model-item label="分类名称" prop="name">
                    <a-input v-model="newCate.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑分类区域 -->
        <a-modal title="编辑分类"
                :visible="editCateVisible"
                @ok="editCateOk"
                @cancel="editCateCancel">
            <a-form-model :model="CateInfo" :rules="CateRules" ref="addCateRef">
                <a-form-model-item label="分类名称" prop="name">
                    <a-input v-model="CateInfo.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'id',
        key:'id',
        width: '10%',
        align: 'center'
    },
    {
        title: '分类名称',
        dataIndex: 'name',
        key:'Catename',
        width: '20%',
        align: 'center'
    },
    {
        title: '操作',
        key:'action',
        width: '30%',
        scopedSlots: {customRender:'action'},
        align: 'center'
    }
]

export default {
    data() {
        return {
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                pageSize: 5,
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
            },
            Catelist: [],
            columns,
            queryParam: {
                pagesize: 5,
                pagenum: 1,
            },
            visible: false,

            CateInfo: {
                id:0,
                name:'',
            },
            newCate: {
                name:'',
            },
            addCateVisible: false,
            CateRules:{
                name:[
                        {
                            validator: (rule, value, callback) => {
                                if (this.CateInfo.name == '') {
                                    callback(new Error("请输入分类名"))
                                } else {
                                    callback()
                                }
                            }, trigger: 'blur',
                        }],
            },
            addCateRules:{
                name:[
                        {
                            validator: (rule, value, callback) => {
                                if (this.newCate.name == '') {
                                    callback(new Error("请输入分类名"))
                                } else {
                                    callback()
                                }
                            }, trigger: 'blur',
                        }],
            },
            editCateVisible: false,
        }
    },
    created(){
        this.getCateList()
    },
    methods: {
        // 获取分类列表
        async getCateList() {
            const { data: res } = await this.$http.get('category', {
                params: {
                    pageSize: this.queryParam.pagesize,
                    pageNum: this.queryParam.pagenum,
                    },
            })
            
            if (res.status != 200) return this.$message.error(res.message)
            this.Catelist = res.data
            this.pagination.total = res.total
        },
    
        handleTableChange(pagination, filters, sorter) {
            var pager = { ...this.pagination }
            pager.current = pagination.current
            pager.pageSize = pagination.pageSize
            this.queryParam.pagenum = pagination.pageSize
            this.queryParam.pagenum = pagination.current

            if (pagination.pageSize !== this.pagination.pageSize) {
                this.queryParam.pagenum = 1
                pager.current = 1
            }
            this.pagination = pager
            this.getCateList()
        },

        // 删除分类
        async deleteCate(id) {
            this.$confirm({
                title: '提示： 请再次确认',
                content: '一旦删除将无法恢复！',
                onOk: async () => {
                    const res = await this.$http.delete(`category/${id}`)
                    if (res.status != 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getCateList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });
        },

        // 新增分类
        async addCateOk() {
            this.$refs.addCateRef.validate(async (valid)=> {
                if (!valid) return this.$message.error("参数不符合要求，请重新输入")
                const { data: res } = await this.$http.post('category/add', {
                    name: this.newCate.name,
                    password: this.newCate.password,
                    role: this.newCate.role,
                })
                if (res.status != 200) return this.$message.error(res.message)
                this.$message.success('添加分类成功')
                this.addCateVisible = false
                await this.getCateList()
            })
        },

        // 取消新增分类
        addCateCancel() {
            this.$refs.addCateRef.resetFields()
            this.addCateVisible = false
            this.$message.info('添加已取消')
        },

        // 编辑分类
        async editCate(id) {
            this.editCateVisible = true
            const { data: res} = await this.$http.get(`category/${id}`)
            this.CateInfo = res.data
            this.CateInfo.id = id
        },
        editCateOk() {
            this.$refs.addCateRef.validate(async (valid)=> {
                if (!valid) return this.$message.error("参数不符合要求，请重新输入")
                const { data: res } = await this.$http.put(`category/${this.CateInfo.id}`, {
                    name: this.CateInfo.name,
                    role: this.CateInfo.role,
                })
                if (res.status != 200) return this.$message.error(res.message)
                this.$message.success('编辑分类成功')
                this.editCateVisible = false
                this.getCateList()
            })
        },
        editCateCancel() {
            this.$refs.addCateRef.resetFields()
            this.editCateVisible = false
            this.$message.info('编辑已取消')
        }
    },
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>