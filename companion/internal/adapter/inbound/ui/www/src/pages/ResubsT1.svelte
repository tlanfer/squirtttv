<script>
    import EventPage from "../components/EventPage.svelte";
    import {config} from "../lib/config.js";
    import EventItem from "../components/EventItem.svelte";
    import EventEditor from "../components/EventEditor.svelte";

    const unit = "month";
    const units = "months";

    $: events = ($config.events?.resubt1 || []).sort((a, b) => a.amount - b.amount);

    function addEvent() {
        const all = [{amount: 0}, ...events].map(e => e.amount).sort(a => a.amount)
        const amount = all[all.length - 1] + 1;
        $config.events.resubt1 = [...events, {
            amount: amount,
            match: "minimum",
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
        $config.events.resubt1[selectedIndex] = selectedEvent
        showModal = false;
        selectedIndex = -1;
        selectedEvent = {};
    }

    function del() {
        console.log("delete", selectedIndex, selectedEvent)
        $config.events.resubt1 = events.filter((e, i) => i !== selectedIndex);
        showModal = false;
        selectedIndex = -1;
        selectedEvent = {};
    }

</script>


<EventPage {addEvent}>
    {#each events as event, i}
        <EventItem unit={unit} units={units} event={event} onSelect={()=>{ selectEvent(i) }}/>
    {/each}
</EventPage>

<EventEditor bind:showModal units={units} event={selectedEvent} onSave={save} onDelete={del}/>