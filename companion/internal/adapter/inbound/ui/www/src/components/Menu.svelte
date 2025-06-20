<script>
import SaveState from "./SaveState.svelte";
import IntegrationState from "./IntegrationState.svelte";
import Version from "./Version.svelte";
</script>

<h1>Squirtttv</h1>

<h2>Global</h2>
<ul>
    <li><a href="/settings">Settings</a></li>
    <li><a href="/devices">Devices</a></li>
</ul>

<h2>Events</h2>
<ul>
    <li><a href="/donations">Donations</a></li>
    <li><a href="/bits">Bits</a></li>
    <li><a href="/gifts">Gift subs</a></li>
    <li><a href="/resubs1">Resubs <span class="sub">Tier 1</span></a></li>
    <li><a href="/resubs2">Resubs <span class="sub">Tier 2</span></a></li>
    <li><a href="/resubs3">Resubs <span class="sub">Tier 3</span></a></li>
    <li><a href="/chat">Chat</a></li>
</ul>

<h2>Status</h2>
<div id="status">
    <SaveState />
    <IntegrationState />

    <p>
        <Version />
    </p>

    <div class="hint">
        <b>Hint:</b>
        You can close this window when you are done configuring your Squirtttv.
        Just keep the companion app running in the background.
    </div>
</div>

<style>
    h1 {
        text-align: center;
        font-weight: 900;
        font-size: 4em;
        margin: 0;
    }
    h2 {
        margin: 1em 0 .2em 0;
    }

    ul {
        display: flex;
        flex-direction: column;
        list-style: none;
        padding: 2px;
        gap: 5px;
        margin: 0;
    }

    li {
        font-size: 1.2em;
        margin: 3px 0;
    }

    a {
        color: inherit;
        text-decoration: none;
    }

    a .sub {
        color: #999;
        font-size: 0.8em;
    }

    a:hover {
        border-bottom: 2px solid tomato;
    }

    #status {
        font-size: 0.8em;

    }

    .hint {
        margin-top: 1em;
        color: #999;
    }
</style>
