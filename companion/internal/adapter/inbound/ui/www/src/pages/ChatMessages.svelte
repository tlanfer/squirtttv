<script>
    import EventPage from "../components/EventPage.svelte";
    import {config} from "../lib/config.js";
    import ChatItem from "../components/ChatItem.svelte";
    import ChatEditor from "../components/ChatEditor.svelte";

    $: events = ($config.events?.chat || []).sort((a, b) => a.trigger - b.trigger);

    function addEvent() {
        $config.events.chat = [...events, {
            trigger: "!squirt",
            devices: [],
            choose: "all",
            pattern: "2s"
        }];
    }

    let showModal = false;
    let selectedEvent = {};
    let selectedIndex = -1;

    function selectEvent(i) {
        selectedIndex = i;
        selectedEvent = structuredClone(events[i]);
        showModal = true;
    }

    function save() {
        $config.events.chat[selectedIndex] = selectedEvent
        showModal = false;
        selectedIndex = -1;
        selectedEvent = {};
    }

    function del() {
        console.log("delete", selectedIndex, selectedEvent)
        $config.events.chat = events.filter((_, i) => i !== selectedIndex);
        showModal = false;
        selectedIndex = -1;
        selectedEvent = {};
    }

</script>


<EventPage {addEvent} helpText="The first matching trigger will be used. This may be random.">
    {#each events as event, i}
        <ChatItem event={event} onSelect={()=>{ selectEvent(i) }}/>
    {/each}
</EventPage>

<ChatEditor bind:showModal event={selectedEvent} onSave={save} onDelete={del}/>