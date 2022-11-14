import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import Index from '../components/admin/Index.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'
import Profile from '../components/user/Profile.vue'
import CommentList from '../components/comment/commentList.vue'


// 路由重复点击捕获错误
// const originalPush = VueRouter.prototype.push
// VueRouter.prototype.push = function push(location, onResolve, onReject) {
//   if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
//   return originalPush.call(this, location).catch(err => err)
// }

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/',
    name: 'Admin',
    component: Admin,
    children: [
      {path: 'index', component: Index},
      {path: 'addart', component: AddArt},
      {path: 'addart/:id', component: AddArt, props:true},
      {path: 'artlist', component: ArtList},
      {path: 'catelist', component: CateList},
      {path: 'userlist', component: UserList},
      {path: 'profile', component: Profile, meta: {title: '个人设置'}},
      {
        path: 'commentlist',
        component: CommentList,
        meta: {
          title: '评论管理'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  // to and from are both route objects. must call `next`.
  const token = window.sessionStorage.getItem('token')
  if (to.path == '/login') return next()
  if (!token){
    next('/login')
  } else {
    next()
  }
})

export default router
