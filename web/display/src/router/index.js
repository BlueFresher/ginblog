import Vue from 'vue'
import VueRouter from 'vue-router'

import HomeView from '../views/HomeView/HomeView.vue'

//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/FirstView/FirstView.vue'),
    meta: {
      title: '首页'
    }
  },
  {
    path: '/blog',
    name: 'blog',
    component: () => import('../views/BlogView/BlogView.vue'),
    meta: {
      title: '博客'
    },
    children:[
      {path: '/', component: () => import('../views/BlogView/components/ArticleList.vue'), meta:{title: '欢迎来到GinBlog'}},
      {path: '/article/detail/:id', component: () => import('../views/BlogView/components/Detail.vue'), meta:{title:'文章详情'}, props:true},
      {path: '/category/:cid', component: () => import('../views/BlogView/components/CateList.vue'), meta: { title: '分类信息' }, props: true},
      {path: '/search/:title', component: () => import('../views/BlogView/components/Search.vue'), meta: { title: '搜索结果' }, props: true},
    ]
  },
  {
    path: '/back',
    name: 'back',
    component: () => import('../views/BackView/BackView.vue'),
    meta: {
      title: '后台'
    }
  },
  {
    path: '/404',
    name: '404',
    component: () => import('../views/404View.vue'),
    meta: {
      title: '404'
    }
  },

  { path: '*', redirect: '/404', hidden: true }
]

const router = new VueRouter({
  routes
})

export default router
