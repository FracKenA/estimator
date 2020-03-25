new Vue({
  el: '#app',
  data: {
    showNav: false
  },
  delimiters: ['{%', '%}']
});

new Vue({
  el: '#estimator',
  data: {
    // Assumptions Consts
    procLogfilesPerDay: 40,
    gwInstallsPerDay: 1,
    l3ServersPerDay: 1.5,
    l4ServersPerDay:  1,
    l5ServersPerDay:  0.75,
    dashboardLower: 10,
    dashboardUpper: 15,
    documentationPercent: 15,
    ktPercent: 5,
    contingencyPercent: 10,
    l1l2LowerLimitDays: 10,

    // System Assumptions 
    numEnv: 1,
    gwPerTeam: 1,
    numProcPerServer: 5,
    numLogPerServer: 5,
    percentL3:  30,
    percentL4: 20,
    percentL5: 10,


    // Project Parameters 
    // with resonable defaults 
    gwNum: 2,
    numServers: 20
  },
  computed: {
    // Test vars
    test: function () {
      return this.numEnv * 2 
    },
    numl1l2ServersPerDay: function () {
      return this.procLogfilesPerDay/(this.numProcPerServer + this.numLogPerServer)
    },
    totalLogProc: function () {
      return this.numServers * (this.numProcPerServer + this.numLogPerServer)
    },
    numServersLower: function () {
      return this.numl1l2ServersPerDay * this.l1l2LowerLimitDays
    },
    // Estimate table vars
    l1l2EstimateLower: function () {
      if (this.numServers > this.numServersLower) 
	return 10
      else
	return this.numServers/this.numl1l2ServersPerDay 
    },
    l1l2EstimateUpper: function () {
      if (this.numServers > this.numServersLower) 
	return 20
      else
	// TODO change this from NA
	// return this.numServers/this.numl1l2ServersPerDay 
	return "N/A"
    }
  },
  delimiters: ['{%', '%}']
});
