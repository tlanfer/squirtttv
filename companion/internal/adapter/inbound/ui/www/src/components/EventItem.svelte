<script>
    import { scale } from 'svelte/transition';
    import EventEditor from "./EventEditor.svelte";

    export let onSelect = () => {
    };
    export let unit;
    export let units;

    export let event;
    let showModal = false;
</script>

<EventEditor bind:showModal units={units} event={event}/>

<div class="item" transition:scale={{duration: 100}}>
    <div class="value">
        <span class="number">{event.amount}<small>{event.match === "minimum" ? "+" : ""}</small></span>
        <span class="unit">{event.amount === 1 ? unit : units}</span>
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
        width: 100px;
        display: flex;
        flex-direction: row;
        gap: 0.5em;
        align-items: baseline;
        margin-bottom: 0.5em;
    }

    .value .number {
        font-size: 2em;
        font-weight: 900;
        width: 100%;
        text-align: center;
        border: none;
        font-family: inherit;
        box-sizing: border-box;
    }
    .value .number small {
        vertical-align: middle;
        font-size: 0.5em;
    }
</style>