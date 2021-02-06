<template>
  <div>
    <div>
      <input v-model="branch" placeholder="branch" />
      <button @click="createBranch">Create New Branch</button>
    </div>
    <h2>Branches</h2>
    <table style="width:100%">
      <tr>
        <th>Name</th>
        <th>Commit</th>
        <th>Protected</th>
      </tr>
      <tr v-for="branch in branches" :key="branch.name">
        <td>{{ branch.name }}</td>
        <td>{{ branch.commit.sha }}</td>
        <td>{{ branch.protected }}</td>
        <td><button @click="commit(branch)">commit on branch</button></td>
        <td>
          <button @click="pullRequest(branch)">Create Pull Request</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "Branch",
  data() {
    return {
      repo: "",
      branches: [],
      branch: ""
    };
  },
  mounted() {
    this.repo = this.$route.params.repo;
    this.getBranches();
  },
  methods: {
    getBranches() {
      axios
        .get("/server/repos/" + this.repo + "/branches")
        .then(response => {
          this.branches = response.data;
        })
        .catch(err => {
          console.log(err);
        })
        .catch(err => console.error(err));
    },
    createBranch() {
      if (this.branch != "" && this.checkBranchUnique()) {
        var request = { repo: this.repo, new_branch: this.branch };
        axios
          .post("/server/branch", request)
          .then(response => {
            this.branch = "";
            this.getBranches();
          })
          .catch(err => {
            console.error(err);
            this.branch = "";
          });
      } else {
        this.branch = "";
        console.log("Please provide unique branch name");
      }
    },
    checkBranchUnique() {
      var branch = this.branch;
      var valObj = this.branches.filter(function(elem) {
        if (elem.name == branch) return elem.name;
      });
      if (valObj.length > 0) {
        return false;
      }
      return true;
    },
    commit(branch) {
      this.$router.push({
        name: "CommitForm",
        query: { repo: this.repo, branch: branch.name }
      });
    },
    pullRequest(branch) {
      this.$router.push({
        name: "PullRequestForm",
        query: { repo: this.repo, branch: branch.name }
      });
    }
  }
};
</script>

<style scoped>
th,
td {
  padding: 15px;
  text-align: left;
}
</style>
