<script>
    import {status} from "../lib/status.js";
</script>

<div class="status" class:connected={$status.twitch}>Twitch { $status.twitch ? "connected" : ""}</div>
<div class="status" class:connected={$status.streamlabs}>Streamlabs { $status.streamlabs ? "connected" : ""}</div>
<div class="status" class:connected={$status.streamelements}>Streamelements { $status.streamelements ? "connected" : ""}</div>

<style>
    .status {
        margin-top: 5px;
    }
    .status:before {
        content: "● ";
        color: tomato;
    }

    .status.connected:before {
        content: "● ";
        color: green;
    }
</style>