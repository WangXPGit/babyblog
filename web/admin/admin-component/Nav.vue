<template>
    <a-layout-sider breakpoint="lg" v-model="collapsed">
        <div class="log">
            <!-- <span>Baby Blog</span> -->
            <span>{{ collapsed ? 'Blog' : 'Baby Blog'}}</span>
        </div>
       
        <a-menu theme="dark" mode="inline">
            <a-menu-item>
                <a-icon type="dashboard"></a-icon>
                <span>仪表盘</span>
            </a-menu-item>
            <a-sub-menu>
                <span slot="title"><a-icon type="file" /><span>文章管理</span></span>
                <a-menu-item><a-icon type="form" /><span>写文章</span></a-menu-item>
                <a-menu-item><a-icon type="snippets" /><span>文章列表</span></a-menu-item>
            </a-sub-menu>
            <a-menu-item><a-icon type="book" /><span>分类列表</span></a-menu-item>
            <a-menu-item><a-icon type="user" /><span>用户列表</span></a-menu-item>
        </a-menu>
    </a-layout-sider>
</template>

<script>
export default {
    data() {
        return {
            collapsed:false
        }
    }
}
</script>
<style scoped>
.log {
    height: 32px;
    margin: 16px;
    background: white;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 17px;
}
</style>