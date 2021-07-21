<template>
    <div>
        <h3>用户列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search v-model="queryParam.username" placeholder="输入用户名查找" enter-button allowClear @search="getUserList" />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible = true">新增</a-button>
                </a-col>
            </a-row>

            <a-table closable destroyOnClose rowKey="username" :columns="columns" :pagination='pagination' :dataSource="userlist" bordered @change="handleTableChange">
                <span slot="role" slot-scope="role">{{ role == 1 ? '管理员' : '用户'}}</span>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right:15px" @click="editUser(data.ID)">编辑</a-button>
                        <a-button type="danger" @click="deleteUser(data.ID)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>

        <!-- 新增用户区域 -->
        <a-modal title="新增用户"
                :visible="addUserVisible"
                @ok="addUserOk"
                @cancel="addUserCancel"
                :destroyOnClose="true">
            <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
                <a-form-model-item label="用户名" prop="username">
                    <a-input v-model="userInfo.username"></a-input>
                </a-form-model-item>

                <a-form-model-item has-feedback label="密码" prop="password">
                    <a-input-password v-model="userInfo.password"></a-input-password>
                </a-form-model-item>

                <a-form-model-item has-feedback label="确认密码" prop="checkpass">
                    <a-input-password v-model="userInfo.checkpass"></a-input-password>
                </a-form-model-item>

                <a-form-model-item label="是否为管理员" prop="role">
                    <a-select defaultValue="2" style="120px" @change="adminChange">
                        <a-select-option key="1" value="1">是</a-select-option>
                        <a-select-option key="2" value="2">否</a-select-option>
                    </a-select>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑用户区域 -->
        <a-modal title="编辑用户"
                :visible="editUserVisible"
                @ok="editUserOk"
                @cancel="editUserCancel">
            <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
                <a-form-model-item label="用户名" prop="username">
                    <a-input v-model="userInfo.username"></a-input>
                </a-form-model-item>
                <a-form-model-item label="是否为管理员" prop="role">
                    <a-select defaultValue="2" style="120px" @change="adminChange">
                        <a-select-option key="1" value="1">是</a-select-option>
                        <a-select-option key="2" value="2">否</a-select-option>
                    </a-select>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        key:'id',
        width: '10%',
        align: 'center'
    },
    {
        title: '用户名',
        dataIndex: 'username',
        key:'username',
        width: '20%',
        align: 'center'
    },
    {
        title: '角色',
        dataIndex: 'role',
        key:'role',
        width: '20%',
        scopedSlots: {customRender:'role'},
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
            userlist: [],
            columns,
            queryParam: {
                username: '',
                pagesize: 5,
                pagenum: 1,
            },
            visible: false,

            userInfo: {
                id:0,
                username:'',
                password:'',
                checkpass:'',
                role: 2,
            },
            addUserVisible: false,
            userRules:{
                username:[
                        {
                            validator: (rule, value, callback) => {
                                if (this.userInfo.username == '') {
                                    callback(new Error("请输入用户名"))
                                }
                                if ([... this.userInfo.username].length < 4 || [... this.userInfo.username].length > 12) {
                                    callback(new Error("用户名应当在4到12个字符之间"))
                                } else {
                                    callback()
                                }
                            }, trigger: 'blur',
                        }],
                checkpass:[
                        {
                            validator: (rule, value, callback) => {
                                if (this.userInfo.checkpass === '') {
                                    callback(new Error("请输入密码"))
                                }
                                if (this.userInfo.password !== this.userInfo.checkpass) {
                                    callback(new Error("密码不一致，请重新输入'"))
                                } else {
                                    callback()
                                }
                            }, trigger: 'blur',
                        }],
                password:[
                        {
                            validator: (rule, value, callback) => {
                                if (this.userInfo.password === '') {
                                    callback(new Error("请输入密码"))
                                }
                                if (this.userInfo.password.length < 6 || this.userInfo.password.length > 20) {
                                    callback(new Error("密码必须为6到20个字符之间"))
                                } else {
                                    callback()
                                }
                            }, trigger: 'blur',
                        }], 
            },
            editUserVisible: false,
        }
    },
    created(){
        this.getUserList()
    },
    methods: {
        async getUserList() {
            const { data: res } = await this.$http.get('users', {
                params: {
                    username: this.queryParam.username,
                    pageSize: this.queryParam.pagesize,
                    pageNum: this.queryParam.pagenum,
                    },
            })
            
            if (res.status != 200) return this.$message.error(res.message)
            this.userlist = res.data
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
            this.getUserList()
        },

        // 删除用户
        async deleteUser(id) {
            this.$confirm({
                title: '提示： 请再次确认',
                content: '一旦删除将无法恢复！',
                onOk: async () => {
                    const res = await this.$http.delete(`user/${id}`)
                    if (res.status != 200) return this.$message.error(res.message)
                    this.$message.success('删除成功')
                    this.getUserList()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                },
            });
        },

        // 新增用户
        async addUserOk() {
            this.$refs.addUserRef.validate(async (valid)=> {
                console.log("abc")
                if (!valid) return this.$message.error("参数不符合要求，请重新输入")
                const { data: res } = await this.$http.post('user/add', {
                    username: this.userInfo.username,
                    password: this.userInfo.password,
                    role: this.userInfo.role,
                })
                if (res.status != 200) return this.$message.error(res.message)
                this.$message.success('添加用户成功')
                this.addUserVisible = false
                this.getUserList()
            })
        },

        // 取消新增用户
        addUserCancel() {
            this.$refs.addUserRef.resetFields()
            this.addUserVisible = false
            this.$message.info('添加已取消')
        },

        adminChange(value) {
            this.userInfo.role = value
            console.log(this.userInfo.role)
        },

        // 编辑用户
        async editUser(id) {
            this.editUserVisible = true
            const { data: res} = await this.$http.get(`user/${id}`)
            this.userInfo = res.data
            this.userInfo.id = id
        },
        editUserOk() {
            this.$refs.addUserRef.validate(async (valid)=> {
                if (!valid) return this.$message.error("参数不符合要求，请重新输入")
                const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
                    username: this.userInfo.username,
                    role: this.userInfo.role,
                })
                if (res.status != 200) return this.$message.error(res.message)
                this.$message.success('编辑用户成功')
                this.editUserVisible = false
                this.getUserList()
            })
        },
        editUserCancel() {
            this.$refs.addUserRef.resetFields()
            this.editUserVisible = false
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