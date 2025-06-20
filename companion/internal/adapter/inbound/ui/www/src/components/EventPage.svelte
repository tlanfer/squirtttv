<script>
    import FancyButton from "./FancyButton.svelte";

    export let addEvent;

    export let helpText = "Add as many events as you want. Exact matching events will be preferred. For events matching 'at least' amounts, the highest amount that still matches will be used.";
</script>

<p class="help">
    {helpText}
</p>

<div id="eventPage">
    <div class="eventList">
        <slot />
        <FancyButton --bg="hsl(100, 50%, 50%)" on:click={addEvent} big>Add event</FancyButton>
    </div>
</div>

<style>
    #eventPage {
        display: flex;
        flex-direction: column;
        gap: 1em;
        justify-content: center;
    }

    .eventList {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
        grid-gap: 1em;
        grid-auto-rows: 100px
    }

</style>