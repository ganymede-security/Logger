<script setup lang="ts">

var d = new Date()
let hr: number; hr = d.getUTCHours();console.log(hr)
let min: number; min = d.getUTCMinutes();console.log(min)




</script>

<template>
    <header>
      <h2>Home</h2>
    </header>
    <fieldset>
        <main>
            <div class="sidenav">
                <a href="#">Average End-to-End Latency </a>
                <a href="#">Average Response Time </a>
                <a href="#">Errors Per Minute </a>
                <a href="#">Exceptions Per Minute </a>
            </div>
        </main>
        <div id="app">
          <h2 class="summary-title">Summary</h2>
          <div class="error"></div>
          <p class="error-text">Error</p>
          <div class="info"></div>
          <p class="info-text">Info</p>
          <canvas ref="log-canvas" width="550" height="450"></canvas>
        </div>
    </fieldset>
</template>

<style scoped>

.sidenav {
  height: 100%;
  width: 320px;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #121212;
  overflow-x: hidden;
  padding-top: 20px;
}

.sidenav a {
  padding: 6px 8px 6px 16px;
  text-decoration: none;
  font-size: 22px;
  color: #818181;
  display: block;
}

.error {
  height: 10px;
  width: 10px;
  background-color: #c91c45;
  border-radius: 50%;
  display: inline-block;
  outline-color: black;
  outline-style: solid;
  outline-width: thin;
}

.info {
  height: 10px;
  width: 10px;
  background-color: #4735b8;
  border-radius: 50%;
  display: inline-block;
  outline-color: black;
  outline-style: solid;
  outline-width: thin;
}

.info-text{
  display:inline-block;
  padding-left: 10px;
}

.error-text{
  display:inline-block;
  padding-left: 10px;
  padding-right: 10px;
}

.summary-title{
  color: white;
  -webkit-text-stroke: 1px;
  -webkit-text-stroke-color: black;
}

</style>
