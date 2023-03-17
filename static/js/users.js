const apiURL = "//randomuser.me/api/?results=10";

const App = new Vue({
  el: "#app",
  data: {
    users: [] },

  created: function () {
    this.fetchData();
  },
  methods: {
    fetchData: function () {
      let self = this;
      axios.
      get(apiURL).
      then(function (response) {
        self.users = response.data.results;
      }).
      catch(function (error) {
        alert("Error cargando la informacion");
      });
    },
    showMore: function (user) {

    } } });