<template>
<div>
 <h2>Repos</h2>
<table style="width:100%">
  <tr>
    <th>Name</th>
    <th>Owner</th>
    <th>Language</th>
    <th>Details</th>
  </tr>
  <tr v-for="repo in repos" :key="repo.id">
    <td>{{repo.name}}</td>
    <td>{{repo.owner.login}}</td>
    <td>{{repo.language}}</td>
    <td>{{repo.branchesUrl}}</td>
    <td><button @click=View(repo)>View</button></td>
  </tr>
</table>
</div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)
export default {
  name: 'Repository',
  data () {
    return {
      repos: []
    }
  },
  mounted() {
    Vue.axios.get("http://localhost:8080/repos").then((response) => {
    this.repos = response.data
})
  },
  methods: {
    View(repo) {
        this.$router.push({name:'Branch', params:{repo:repo.name}})
    }
  }
}
</script>

<style scoped>
table {
th, td {
  padding: 15px;
  text-align: left;
}

}
</style>