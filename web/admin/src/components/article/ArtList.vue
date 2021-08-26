<template>
    <div>
        <h3>文章列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search v-model="queryParam.title" placeholder="输入文章名查找" enter-button allowClear @search="getArtList" />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addArtVisible = true">新增</a-button>
                </a-col>
            </a-row>

            <a-table closable destroyOnClose rowKey="ID" :columns="columns" :pagination='pagination' :dataSource="Artlist" bordered @change="handleTableChange">
                <span slot="img" slot-scope="img">
                    <img :src="img" />
                </span>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button size="small" type="primary" icon="edit" style="margin-right:15px" @click="editArt(data.ID)">编辑</a-button>
                        <a-button size="small" type="danger" icon="delete" @click="deleteArt(data.ID)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        key:'id',
        width: '5%',
        align: 'center'
    },
    {
        title: '分类名',
        dataIndex: 'Category.name',
        key:'name',
        width: '5%',
        align: 'center'
    },
    {
        title: '文章标题',
        dataIndex: 'title',
        key:'title',
        width: '15%',
        align: 'center'
    },
    {
        title: '文章描述',
        dataIndex: 'desc',
        key:'desc',
        width: '20%',
        align: 'center'
    },
    {
        title: '文章缩略图',
        dataIndex: 'img',
        key:'img',
        width: '10%',
        align: 'center'
    },   
    {
        title: '操作',
        key:'action',
        width: '10%',
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
            Artlist: [],
            columns,
            queryParam: {
                title: '',
                pagesize: 5,
                pagenum: 1,
            },
        }
    },
    created(){
        this.getArtList()
    },
    methods: {
        // 获取文章列表
        async getArtList() {
            const { data: res } = await this.$http.get('article', {
                params: {
                    title: this.queryParam.title,
                    pageSize: this.queryParam.pagesize,
                    pageNum: this.queryParam.pagenum,
                    },
            })
            
            if (res.status != 200) return this.$message.error(res.message)
            this.Artlist = res.data
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
            this.getArtList()
        },

        // 删除文章
        async deleteArt(id) {
            this.$confirm({
                title: '提示： 请再次确认',
                content: '一旦删除将无法恢复！',
                onOk: async () => {
                    const res = await this.$http.delete(`article/${id}`)
                    if (res.status != 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getArtList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });
        },
    },
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>