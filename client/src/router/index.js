import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Repository from '@/components/Repository'
import Branch from '@/components/Branch'
import CommitForm from '@/components/CommitForm'
import PullRequestForm from '@/components/PullRequestForm'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/repos',
      name: 'Repository',
      component: Repository
    },
    {
      path: '/branch/:repo',
      name: 'Branch',
      component: Branch
    },
    {
      path: '/commit',
      name: 'CommitForm',
      component: CommitForm
    },
    {
      path: '/pullRequest',
      name: 'PullRequestForm',
      component: PullRequestForm
    }
  ]
})
