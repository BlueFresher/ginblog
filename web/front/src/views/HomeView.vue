<template>
  <v-app app>
    <TopBar></TopBar>
    
    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="12" md="3">
            <Nav></Nav>
          </v-col>

          <v-col cols="12" md="9">
            <v-sheet min-height="80vh" rounded="lg">
              <!-- vue会复用相同组件，对于路由有多个子路由来说，当在子路由来回切换时，
                会导致页面不刷新的问题，这是因为不再执行created和mounted这些钩子函数，
                可以通过watch来监听$route的变化从而加载不同的组件 -->
                <!-- 通过设置key值为$route.path,因为子路由的path不同，
                  从而避免了组件复用，子路由间来回切换时，页面都会重新加载. -->
              <!-- 这个页面下有子页面也就是子路由，所以需要设置router-view -->
              <router-view :key="$route.fullPath"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <Footer></Footer>
  </v-app> 
</template>

<script>
import TopBar from '../components/TopBar'
import Footer from '../components/Footer'
import Nav from '../components/Nav'
  export default {

    components: {
      TopBar, Footer, Nav
    },
  }
</script>
