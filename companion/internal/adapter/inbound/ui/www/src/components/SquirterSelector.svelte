<script>
    import {slide} from "svelte/transition";
    import FancyButton from "./FancyButton.svelte";
    import {squirters} from "../lib/config.js";

    export let onSelect = (selection) => {

    };
    export let close = () => {
    };

    let doSelect = () => {
        onSelect(selection);
    }

    let selection = [];
</script>

<div transition:slide>
    <h3>Select some squirters</h3>

    <ul>
        {#each $squirters as squirter}
            <li>
                <input type="checkbox" bind:group={selection} value={squirter.host} id={squirter.host}>
                <label for={squirter.host}>
                    <span class="material-symbols-outlined">household_supplies</span>
                    {squirter.name}
                </label>
            </li>
        {/each}
    </ul>

    <div id="actions">
        <FancyButton --bg="hsl(  0, 00%, 50%)" on:click={close}>Back</FancyButton>
        <FancyButton --bg="hsl(100, 50%, 50%)" on:click={doSelect} disabled={selection.length === 0}>Add selected</FancyButton>
    </div>
</div>

<style>
    #actions {
        display: flex;
        justify-content: space-between;
    }

    ul {
        padding: 0;
        list-style: none;
        display: flex;
        flex-direction: row;
        gap: 1em;
    }

    input[type="checkbox"] {
        display: none
    }

    label {
        display: flex;
        align-items: center;
        padding: 0.5em;
        border-radius: 5px;
        border: 1px solid var(--fg, black);
        cursor: pointer;
        user-select: none;
    }

    input:checked+label {
        background-color: var(--bg, tomato);
        color: var(--fg, white);
        border-color: white;
    }

</style>