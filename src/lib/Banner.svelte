<script>
</script>

<main>
  <div class="cont">
    <div class="col col-1">
      <h1>Luka Đorđević</h1>
      <h2>Control Systems Engineer</h2>
      <img class="profile_pic" src="src/assets/profile_pic.png" alt="profile" />
    </div>
    <div class="col col-2">
      <p>
        Hello! I am Luka, born and currently living in Serbia. Welcome to my
        website where you can see all the past and on-going projects of mine. I
        studyed Mechanical Engineering with major in Control Systems, but I also
        focus a lot on Embedded system programming, electronis and hardware
        design as well as Web development.
      </p>
    </div>
  </div>
</main>

<style>
  p {
    min-width: 300px;
    text-align: justify;
  }
  .col-1 {
    flex: 1;
  }
  .col-2 {
    flex: 3;
  }
  .cont {
    display: flex;
    flex-direction: row;
    align-items: center;
    text-align: center;
    flex-wrap: wrap;
    flex: 1 3;
    /* background-color:green; */
  }
  .col {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  h1 {
    font-size: 3.2em;
    line-height: 1.1;
    margin: 0px;
  }
  h2 {
    font-size: 1.5em;
    line-height: 1.1;
    margin: 0px;
  }
  img {
    margin: 0.5em;
  }
  .profile_pic {
    width: 200px;
    /* margin-top:2em; */
  }
</style>
