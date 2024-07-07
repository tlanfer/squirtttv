<script>

    import {config} from "../lib/config.js";

    if (!$config.settings) {
        $config.settings = {
            twitch: "",
            streamlabs: "",
            streamelements: ""
        }
    }
</script>

<p class="help">
    This page is where you can configure some global settings for Squirtttv.
    You need to provide the channel you want to monitor, as well as the API tokens for Streamlabs and/or Streamelements.
</p>
<p class="help">
    Squirtttv uses the twitch channel to connect to chat and receive events for subscriptions, gift subs and bit cheers.
    The Streamlabs and Streamelements tokens are used to connect to their socket APIs to receive events for donations and tips.
</p>
<p class="help">
    You can provide all three and you should not get any double events.
</p>

<form>
    <label for="channel">
        Twitch Channel
        <span class="help">
            The name of the twitch channel you want to monitor
        </span>
        <input type="text" id="channel" name="channel" bind:value={$config.settings.twitch}>
    </label>

    <label for="streamlabs">
        Streamlabs Socket API Token
        <span class="help">
            The <b>Socket API Token</b> you get from Streamlabs. Go to <a href="https://streamlabs.com/dashboard#/settings/api-settings" target="_blank">your streamlabs dashboard</a>, in your account settings > API settings > API tokens to get it.
        </span>
        <input type="password" id="streamlabs" name="streamlabs" bind:value={$config.settings.streamlabs}>
    </label>

    <label for="streamelements">
        Streamelements Socket API Token
        <span class="help">
            The <b>JWT Token</b> you get from Streamelements to connect to their socket API. Go to <a href="https://streamelements.com/dashboard/account/channels" target="_blank">your streamelements dashboard</a>, in your account settings > Channels to get it.
        </span>
        <input type="password" id="streamelements" name="streamelements" bind:value={$config.settings.streamelements}>
    </label>

    <label for="baseCurrency">
        Base Currency
        <span class="help">
            The currency you use in streamlabs/streamelements. This will be the currency Squirtttv uses to calculate the value of other currencies.<br />
        </span>
        <select bind:value={$config.settings.baseCurrency}>
            <option value="eur">Euro (eur)</option>
            <option value="usd">US Dollar (usd)</option>
            <option value="gbp">British Pound (gbp)</option>
        </select>
    </label>
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        gap: 2em;
    }

    label {
        display: flex;
        flex-direction: column;
        gap: 0.5em;
    }

    input, select {
        padding: 0.5em;
        border-radius: 5px;
        border: 1px solid hsl(0, 0%, 90%);
        box-sizing: border-box;
        height: 35px;
    }
</style>