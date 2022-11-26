<script>
  import emailIcon from "../assets/envelope-regular.svg";
  import profil_pic_url from "../assets/profile_pic3.jpg";
  // let my_date = new Date("1996.06.19");
  let my_date = new Date(1996,5,19);
  let today = new Date();
  let diff = new Date(today.getTime() - my_date.getTime());
  let diff_years = diff.getUTCFullYear() - 1970;
</script>

<main>
  <div class="cont">
    <div class="col col-1">
      <h1>Luka Đorđević</h1>
      <h2>Control Systems Engineer</h2>
      <img class="profile_pic" src="{profil_pic_url}" alt="profile" />
      <div class="cont">
        <img src={emailIcon} class="logo" alt="Svelte Logo" />
        <a href="mailto: lukadj996@gmail.com">contact me</a>
      </div>
    </div>
    <div class="col col-2">
      <p>
        Hello! My name is Luka, I am {diff_years} years old, born and living in Serbia.
        Welcome to my website, where you can see all the past and ongoing projects of mine.
        I studied Mechanical Engineering with major in Control Systems, but I also focus a lot on Embedded system programming, electronics and hardware design as well as Web development.
      </p>
    </div>
  </div>
</main>

<style>
  .logo {
    width: 24px;
  }
  p {
    min-width: 300px;
    text-align: justify;
  }
  .col-1 {
    flex: 1;
  }
  .col-2 {
    flex: 3;
    /* align-self: end; */
  }
  .cont {
    display: flex;
    flex-direction: row;
    align-items: center;
    text-align: center;
    flex-wrap: wrap;
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
    border-radius: 50%;
    /* margin-top:2em; */
  }
</style>
