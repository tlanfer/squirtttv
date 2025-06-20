<script>
    import { scale } from 'svelte/transition';
    import EventEditor from "./EventEditor.svelte";
    import ChatEditor from "./ChatEditor.svelte";

    export let onSelect = () => {
    };

    export let event;
    let showModal = false;
</script>

<ChatEditor bind:showModal event={event}/>

<div class="item" transition:scale={{duration: 100}}>
    <div class="value">
        { event.trigger }
    </div>
    <button class="edit" on:click={onSelect}>Edit <span class="material-symbols-outlined">edit</span></button>
</div>

<style>
    .item {
        max-width: 200px;
        height: 100px;
        display: flex;
        flex-direction: column;
        padding: 1em;
        align-items: center;
        border: 1px solid hsl(0, 0%, 90%);
        border-radius: 5px;
        box-sizing: border-box;
    }

    button.edit {
        background: none;
        border: none;
        padding: 0;
        height: 24px;
    }

    button.edit span {
        font-size: small;
    }


    .value {
        max-width: 180px;
        margin-bottom: 0.5em;
        font-weight: 900;
        text-align: center;
        border: none;
        font-family: inherit;
        box-sizing: border-box;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        padding: 0.5em;
    }
</style>





