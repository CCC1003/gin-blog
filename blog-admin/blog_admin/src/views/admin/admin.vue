<template>
  <div class="gvb_admin">
    <GVBAside></GVBAside>
    <div class="main">
      <header>
        <div class="left">
          <a-breadcrumb>
            <a-breadcrumb-item>Home</a-breadcrumb-item>
            <a-breadcrumb-item><a href="#">Application Center</a></a-breadcrumb-item>
            <a-breadcrumb-item>An Application</a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        <div class="right">
          <div class="icon_actions">
            <i class="fa fa-home"></i>
            <i v-if = "theme" class="fa fa-sun-o" @click="setTheme"></i>
            <i v-else class="fa fa-moon-o" @click="setTheme"></i>
            <i class="fa fa-arrows-alt"></i>
          </div>
          <div class="avatar">
            <img src="https://i1.hdslb.com/bfs/face/7a753e143a5b3ddda9c37c0916788bd1484c3d72.jpg@150w_150h.jpg" alt="">
          </div>
          <div class="drop_menu">
            <a-dropdown placement="bottomRight">
              <a class="ant-dropdown-link" @click="e => e.preventDefault()">
                cong <a-icon type="down" />
                <i class="fa fa-angle-down"></i>
              </a>
            <template #overlay>
              <a-menu @click="menuClick">
                <a-menu-item key="menuClick">
                  <a href="javascript:;">个人中心</a>
                </a-menu-item>
                <a-menu-item key="my_message">
                  <a href="javascript:;">我的消息</a>
                </a-menu-item>
                <a-menu-item key="article_list">
                  <a href="javascript:;">文章列表</a>
                </a-menu-item>
                <a-menu-item key="logout">
                  <a href="javascript:;">注销退出</a>
                </a-menu-item>
              </a-menu>
              </template>
            </a-dropdown>
          </div>
        </div>
      </header>
      <div class="tabs"></div>
      <main></main>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import GVBAside from "../../components/admin/gvb_aside.vue"
import {ref} from 'vue'
const router = useRouter()
function menuClick({key}){
  if (key === "logout"){
    return
  }
  router.push({
    name:key
  })
}
const theme = ref(true) //true代表白天 false代表黑夜

function setTheme() {
  theme.value = !theme.value
  if(theme.value) {
    //白天
    document.documentElement.classList.remove("dark")
  }else {
    //黑夜
    document.documentElement.classList.add("dark")
  }
}
</script>

<style lang="scss">
.gvb_admin {
  width: 100%;
  display: flex;
}



.main {
  width: calc(100% - 240px);

  header {
    height: 60px;
    background-color: white;
    padding: 0 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .right {
      display: flex;
      align-items: center;
    }

    .icon_actions {
      margin-right: 20px;

      i {
        margin-left: 10px;
        cursor: pointer;
        font-size: 16px;
        color: var(--text);
      }

      i:hover {
        color: var(--active)
      }
    }

    .avatar {
      img {
        width: 40px;
        height: 40px;
        border-radius: 50%;
      }
    }
    .drop_menu {
      margin-left: 10px;
    }
  }

  .tabs {
    height: 30px;
    border: 1px solid #f0eeee;
  }

  main {
    background-color: var(--bg);
    height: calc(100% - 90px);
  }
}</style>
