<template>
    <a-layout class="container">
      <a-layout-sider>
          <div class="log">
              <span>Baby Blog</span>
          </div>
          <Nav></Nav>
      </a-layout-sider>
      <a-layout>
        <a-layout-header class="headerBtn">
            <Header></Header>
        </a-layout-header>
        <a-layout-content>
            <router-view></router-view>
        </a-layout-content>
        <a-layout-footer>
            <Footer></Footer>
        </a-layout-footer>
      </a-layout>
    </a-layout>
</template>

<script>
import Nav from '../../admin-component/Nav'
import Footer from '../../admin-component/Footer'
import Header from '../../admin-component/Header'
export default {
    components : {Nav, Footer, Header},
    methods: {
        logout() {
            sessionStorage.clear('token')
            this.$router. push('/login')
        }
    }
}
</script>
<style scoped>
.container {
    height: 100%;
}

.headerBtn {
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

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