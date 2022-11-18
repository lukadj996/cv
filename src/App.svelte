<script>
  import svelteLogo from "./assets/svelte.svg";
  import Project from "./lib/Project.svelte";
  // @ts-ignore
  import Banner from "./lib/Banner.svelte";

  let projects = [
    {
      name: "EuroBot Competition",
      stack: [{val:"C/C++",info:"Low level PID controllers, encoders, sensors, etc."},{val:"Matlab",info:'State maching code generation'},
      {val:"Simulink",info:'Modeling and simulating, tunning PID parameters, etc.'},
       {val:"Python",info:"High level control on Raspberry pi"}],
      desc: "EuroBot is competition that is making autonomous mobile robots...Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.",
      imgs: ["src/assets/img1.jpg", "src/assets/img2.jpg"],
      isOpen:true
    },
    {
      name: "Browser tester",
      stack: [{val:"HTML+JS+CSS",info:'Vanilla page with minnimum JS for buttons, popups, etc.'}, {val:"NodeJS",info:"BackEnd code for testing cookies, authorisation, etc."}, {val:"GCP",info:"AppEngine hosting, SSL cert, load-balancing"}],
      desc: "This simple website was prof of concept for manual and automatic browser capatibilities testing. I developed this app to make browser testing for me and my colegues easyerr. After I presented the website to my project lead we made completly new browser testing website based on this one and added a lot more capabitilyes but it is closed souce and used only in my company",
      imgs: ["src/assets/img3.jpg"],
    },
    {
      name: "Solar energy prject",
      stack: [{val:"GO"}, {val:"GCP"}, {val:"NodeJS"}],
      desc: "EuroBot is competition that is making autonomous mobile robots...Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.",
      imgs: [
        "src/assets/img5.jpg",
        "src/assets/img2.jpg",
        "src/assets/img3.jpg",
        "src/assets/img4.jpg",
      ],
    },{
      name: "Bla Bla",
      stack: [{val:"HTML+JS+CSS",info:'Vanilla page with minnimum JS for buttons, popups, etc.'}, {val:"NodeJS",info:"BackEnd code for testing cookies, authorisation, etc."}, {val:"GCP",info:"AppEngine hosting, SSL cert, load-balancing"}],
      desc: "This simple website was prof of concept for manual and automatic browser ",
      imgs: ["src/assets/img2.jpg"],
    },
  ];
</script>

<main>
  <Banner />
  {#each projects as project}
    <Project {project} />
  {/each}
  <div class="made-with">
     <div>Made with Svelte and GCP with &#9825;</div>
    <div>
      <a href="https://svelte.dev" target="_blank" rel="noreferrer"> 
        <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
      </a>
    </div>
  </div>
</main>

<style>
  .made-with{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
