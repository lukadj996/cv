<script>
  import svelteLogo from "./assets/svelte.svg";
  import Banner from "./lib/Banner.svelte";
  import Project from "./lib/Project.svelte";

  let projects = [
    {
      name: "EuroBot Competition",
      stack: ["C/C++", "Matlab", "Python"],
      desc: "EuroBot is competition that is making autonomous mobile robots...Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.",
      imgs: ["src/assets/img1.jpg", "src/assets/img2.jpg"],
    },
    {
      name: "Browser tester",
      stack: ["Javascript", "NodeJS", "GCP"],
      desc: "This simple website was prof of concept for manual and automatic browser capatibilities testing. I developed this app to make browser testing for me and my colegues easyerr. After I presented the website to my project lead we made completly new browser testing website based on this one and added a lot more capabitilyes but it is closed souce and used only in my company",
      imgs: ["src/assets/img3.jpg"],
    },
    {
      name: "Solar energy prject",
      stack: ["GO", "GCP", "NodeJS"],
      desc: "EuroBot is competition that is making autonomous mobile robots...Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.",
      imgs: [
        "src/assets/img5.jpg",
        "src/assets/img2.jpg",
        "src/assets/img3.jpg",
        "src/assets/img4.jpg",
      ],
    },
  ];
</script>

<main>
  <Banner />
  {#each projects as project}
    <Project {project} />
  {/each}
  <!-- <p class="read-the-docs">
    Made with Svelte
  </p>
  <div>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer"> 
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div> -->
</main>

<style>
</style>
