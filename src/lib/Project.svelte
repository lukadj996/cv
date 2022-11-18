<script>
  export let project;
  import Carousel from "svelte-carousel";
  import { fade,fly,scale,slide,blur,draw } from "svelte/transition";
  let isOpen = project.isOpen;
  function toggle(){
    isOpen = !isOpen
  }
  function in_end(e){
     e.target.parentElement.scrollIntoView()
    }
</script>

<main class='sve'>
  <button on:click={toggle}>{#if isOpen}&#9650{:else}&#9660{/if}{project.name}</button
  >
  {#if isOpen}
    <div transition:blur={{ duration: 300 }} on:introstart={in_end} class="cont">
      <div class="col col-1">
        <div class="wrap">
          <Carousel dots={false}>
            {#each project.imgs as src}
              <img {src} alt="" />
            {/each}
          </Carousel>
        </div>
      </div>
      <div class="col col-2">
        <p class="desc">
          {project.desc}
        </p>
        <div class='ts'>Technologies Summary:</div>
        <table>
          {#each project.stack as stack}
          <tr>
          <td>{stack.val}</td>
          <td>&#10132;</td>
          <td>
          {stack.info}
          </td>
        </tr>
          {/each}
        </table>
      </div>
    </div>
  {/if}
</main>

<style>
  table td:first-of-type{
    text-align: right;
    color:darkorange;
  }
  table td:last-of-type{
    text-align: left;
  }
  .ts{
    padding: 0px;
    font-weight: bold;
    color:darkorange;
  }
  .desc {
    min-width: 260px;
    text-align: justify;
    padding-right: 0.5em;
    padding-left: 0.5em;
  }
  .wrap {
    max-width: 400px;
    box-sizing: content-box;
    width: 100%;
  }
  button {
    border: none;
    background: none;
    display: block;
    
    font-size: 1.5em;
    cursor: pointer;
    margin: 0;
    width: 100%;
    padding: 0.2em 0px;
    text-align: left;
    -webkit-tap-highlight-color: transparent;
  }
  button:focus {
    outline: 0;
  }
  .col-1 {
    flex: 10;
    box-sizing: content-box;
    width: 100%;
  }
  .col-2 {
    flex: 50;
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
  main {
    background-color: rgb(89, 89, 89);
    border-radius: 0.5em;
    margin: 0.5em 0px;
    padding: 0.5em 0px;
  }
</style>
