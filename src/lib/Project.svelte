<script>
  export let project;
  import Carousel from "svelte-carousel";

  import { fade } from "svelte/transition";

  let isOpen = false;
  const toggle = () => (isOpen = !isOpen);
</script>

<main>
  <button on:click={toggle}
    ><svg
      style="tran"
      width="20"
      height="20"
      fill="none"
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path d="M9 5l7 7-7 7" />
    </svg>{project.name}</button
  >
  {#if isOpen}
    <div transition:fade={{ duration: 300 }} class="cont">
      <div class="col" />
      <Carousel dots={false}>
        {#each project.imgs as src}
          <img {src} alt="" />
        {/each}
      </Carousel>
    </div>
  {/if}
</main>

<style>
  button {
    border: none;
    background: none;
    display: block;
    color: inherit;
    font-size: 1.5em;
    cursor: pointer;
    margin: 0;
    width: 100%;
    padding-bottom: 0.3em;
    padding-top: 0.3em;
    text-align: left;
  }
  button:focus {
    outline: 0;
  }
  .cont {
    width: 300px;
    /* background-color: lightgrey; */
  }

  main {
    background-color: rgb(89, 89, 89);
    border-radius: 0.5em;
    /* padding: 1em; */
    margin: 0.5em;
  }
</style>
