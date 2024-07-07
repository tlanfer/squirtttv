<script>
    import Description from "../components/Description.svelte";
    import ActionButton from "../components/ActionButton.svelte";

    import { config } from "../lib/config.js";
    import Modal from "../components/Modal.svelte";
    import SquirterFinder from "../components/SquirterFinder.svelte";
    import FancyButton from "../components/FancyButton.svelte";

    $: devices = $config.devices || [];

    let showModal = false;
    let selected = -1;
    let selectedName = "";

    let addDevice = (newDevice) => {
        config.update(c => {
            c.devices = c.devices || [];

            c.devices.push({name: `New squirter ${c.devices.length+1}`, host: newDevice});

            return c;
        });
    };

    let removeDevice = index => () => {
        config.update(c => {
            c.devices.splice(index, 1);
            return c;
        });
    };

    let renameDevice = index => () => {
        selected = index;
        showModal = true;
        selectedName = devices[index].name;
    }

    $: if (!showModal && selected!==-1) {
        $config.devices[selected].name = document.querySelector("input").value;
        selected = -1;
    }

    $: if(showModal){
        setTimeout(()=>document.querySelector("input").focus(), 200);
    }
    let showFinder = false;
</script>

<p class="help">
    In this menu you add any devices you wish to use. You can connect as many as you want.<br/>
    If you have more than one device, you should give them names, so you can identify them when configuring events.
</p>

<Modal bind:showModal>
    <h3>Rename device</h3>
    <form on:submit|preventDefault={()=>showModal=false}>
        <input type="text" placeholder="New name" bind:value={selectedName}/>
    </form>
</Modal>

<SquirterFinder bind:showModal={showFinder} addSquirter={addDevice}/>

<div id="deviceList">
    {#each devices as device, index}
        <div class="device">
            <span class="material-symbols-outlined">household_supplies</span>
            <div class="info">
                <div class="name">{device.name}</div>
                <div class="address">{device.host}</div>
            </div>
            <div class="actions">
                <button class="btn-text" on:click={renameDevice(index)}>Rename</button>
                <button class="btn-text" on:click={removeDevice(index)}>Delete</button>
            </div>
        </div>
    {:else}
        <li>No devices added yet</li>
    {/each}
    <div class="main_action">
        <FancyButton --bg="hsl(100, 50%, 50%)" big on:click={()=>{showFinder = true}}>Add squirter</FancyButton>
    </div>
</div>

<style>
    div#deviceList {
        list-style: none;
        padding: 0;
        display: flex;
        flex-direction: column;
        gap: 1em;
    }

    div.device {
        height: 100px;
        border: 1px solid hsl(0, 0%, 90%);
        border-radius: 5px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        font-weight: bold;
    }

    .info {
        flex-grow: 1;
    }

    .material-symbols-outlined {
        font-size: 3em;
        width: 100px;
        text-align: center;
    }

    .name, .address {
        margin: 10px 0;
        font-size: 0.8em;
    }


    button.btn-text {
        background: none;
        border: none;
        border-bottom: 2px solid transparent;
        margin: 5px 0;
    }

    button.btn-text:hover {
        border-bottom: 2px solid tomato;
    }

    .actions {
        margin-right: 1em;
        display: flex;
        flex-direction: column;
    }

    .main_action {
        text-align: center;
    }

    h3 {
       margin-top: 0;
    }

    input[type="text"] {
        width: 100%;
        padding: 10px;
        border: 1px solid hsl(0, 0%, 90%);
        border-radius: 5px;
        box-sizing: border-box;
    }
</style>