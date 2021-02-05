<template>
  <div class="hello">
        Title : <input v-model="title" placeholder="title">
        Description : <input v-model="description" placeholder="description">
        <button @click=createCommit> create </button>
  </div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)
export default {
  name: 'PullRequestForm',
  data () {
    return {
        description: "",
        repo:"",
        branch:"",
        title:""
    }
  },
  mounted() {
      this.repo = this.$route.query.repo
      this.branch = this.$route.query.branch
  },
  methods: {
      createCommit() {
        var files = []
        files.push(this.file)
        var request = {"repo":this.repo, "new_branch": this.branch, "description":this.description,"title":this.title}
        Vue.axios.post("http://localhost:8080/pullRequest", request).then((response) => {
          alert("Successfully created pull request")
          this.title = ""
          this.description = ""
          this.$router.push({name:'Branch', params:{repo:this.repo}})
        })
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>