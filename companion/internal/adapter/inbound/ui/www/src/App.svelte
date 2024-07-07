<script>
  import router from "page";
  import routes from './routes.js'
  import Menu from "./components/Menu.svelte";

  let page;
  let title;
  let params;

  routes.forEach(route => {
    router(
      route.path,
      (ctx, next) => {
        params = ctx.params;
        next();
      },
      () => {
        page = route.component;
        title = route.title;
      }
    )
  })

  router.start();
</script>

<div id="app">
  <nav>
    <Menu />
  </nav>

  <main>
    <h1>{title}</h1>
    <svelte:component this={page} />
  </main>
</div>

<footer>
  Made by Tlanfer 😎, check <a href="https://github.com/tlanfer/squirtttv" target="_blank">Github</a> for updates.
</footer>

<style>
  #app {
    box-sizing: border-box;
    padding: 4em 4em 1em;

    display: flex;
    flex-direction: row;
    justify-content: stretch;
    gap: 3em;

    max-width: 1400px;
    margin: 0 auto;
  }

  nav, main {
    padding: 1em;
    background-color: white;
    border-radius: 10px;
    box-shadow: 0 0 5px hsla(0, 0%, 0%, 0.05);
  }

  nav {
    flex-basis: 15%;
  }

  main {
    flex-grow: 1;
    overflow-y: auto;
  }

  main h1 {
    line-height: 2em;
  }

  footer, footer a {
    text-align: center;
    margin-top: 1em;
    color: hsla(0, 0%, 0%, 0.3);
  }


</style>
