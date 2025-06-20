<script>
    import {slide} from "svelte/transition";
    import Modal from "./Modal.svelte";
    import SquirterSelector from "./SquirterSelector.svelte";
    import FancyButton from "./FancyButton.svelte";
    import {squirters} from "../lib/config.js";
    import {test} from "../lib/test.js";
    import Gifts from "../pages/Gifts.svelte";

    export let onSave = () => {};
    export let onDelete = () => {};

    export let showModal = false;
    export let event = {};

    let selectorActive = false;

    function unique(value, index, array) {
        return array.indexOf(value) === index;
    }

    let onSquirterSelect = (selectedSquirters) => {
        selectorActive = false;
        event.devices = [ ...event.squirters || [], ...selectedSquirters].filter(unique)
    }

    $: if(!showModal) {
        selectorActive = false;
    }

    function squirterName(squirter) {
        return $squirters.find(s => s.host === squirter).name;
    }

    function removeSquirter(squirter) {
        event.devices = event.devices.filter(s => s !== squirter);
    }

    function onTest() {
        test(event.pattern, event.choose, event.devices);
    }

    $: configValid = event.trigger && event.pattern && event.choose && (event.choose === "all" || event.devices);
</script>

<Modal bind:showModal>
    {#if (selectorActive)}
        <SquirterSelector onSelect={onSquirterSelect} close={()=>{selectorActive=false}}/>
    {:else}
        <div class="config" transition:slide>
            <div class="amount">
                <h3>Trigger</h3>
                <div class="help">
                    When this term appears in a chat message, the event will be triggered.
                    Select all roles that are allowed to trigger this event.
                </div>
                <input type="text" class="trigger_input" bind:value={event.trigger}>
                <p class="roles">
                    <label><input type="checkbox" name="role" bind:group={event.roles} value="mod">Moderators</label>
                    <label><input type="checkbox" name="role" bind:group={event.roles} value="sub">Subscribers</label>
                    <label><input type="checkbox" name="role" bind:group={event.roles} value="pleb">Plebs</label>
                </p>
            </div>
            <div class="pattern">
                <h3>Pattern</h3>
                <div class="help">
                    A list of comma-separated millisecond durations of squirts or pauses. <br />
                    For example, a pattern of 1000, 300, 1000 will squirt for 1s, pause for 300ms, and then squirt for another 1s.
                </div>
                <input type="text" class="pattern_input" bind:value={event.pattern}>
            </div>
            <div>
                <h3>Squirters</h3>
                <div class="help">
                    Select all squirters you want to enable for this event
                </div>
                <div class="squirter_list">
                    <label><input type="radio" name="squirter_mode" bind:group={event.choose} value="all">All</label>
                    <label><input type="radio" name="squirter_mode" bind:group={event.choose} value="allOf">All of</label>
                    <label><input type="radio" name="squirter_mode" bind:group={event.choose} value="oneOf">One of</label>

                    {#if event.choose !== "all"}
                        {#if event.devices}
                            {#each event.devices as device}
                                <FancyButton --bg="hsl(100, 0%, 50%)" on:click={()=>removeSquirter(device)}>
                                    {squirterName(device)}
                                    <span class="material-symbols-outlined btn-text">close_small</span>
                                </FancyButton>
                            {/each}
                        {/if}
                        <FancyButton --bg="hsl(100, 50%, 50%)" on:click={()=>{selectorActive=true}}><span class="material-symbols-outlined">add</span></FancyButton>
                    {/if}
                </div>
            </div>
            <div class="actions">
                <FancyButton --bg="hsl(0, 50%, 50%)" on:click={onDelete}>Delete</FancyButton>
                <FancyButton --bg="hsl(200, 50%, 50%)" on:click={onTest} disabled={!configValid}>Test</FancyButton>
                <FancyButton --bg="hsl(100, 50%, 50%)" on:click={onSave} disabled={!configValid}>Save</FancyButton>
            </div>
        </div>
    {/if}
</Modal>

<style>
    .material-symbols-outlined {
        font-size: 1.3em;
        vertical-align: -3px;
    }
    .config {
        display: flex;
        flex-direction: column;
        gap: 1em;
    }

    .config h3 {
        margin: 0 0 0.5em;
        font-size: 1em;
    }

    .pattern {
        display: flex;
        flex-direction: column;
        gap: 0.5em;
    }

    input[type="text"] {
        width: 100%;
        padding: 0.5em;
        border-radius: 5px;
        border: 1px solid hsl(0, 0%, 90%);
        box-sizing: border-box;
        height: 35px;
    }

    .roles label {
        display: inline-block;
        margin-right: 1em;
    }

    input[type="checkbox"] {
        margin-right: 0.5em;
    }

    .squirter_list {
        display: flex;
        gap: 0.5em;
        flex-wrap: wrap;
    }

    .squirter_list label {
        line-height: 35px;
        vertical-align: middle;
    }

    .actions {
        display: flex;
        gap: 0.5em;
        justify-content: space-between;
    }

    .btn-text {
        background: none;
        border: none;
        color: white;
        padding: 0;
    }

</style>