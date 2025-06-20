import Devices from "./pages/Devices.svelte";
import Settings from "./pages/Settings.svelte";
import Donations from "./pages/Donations.svelte";
import Bits from "./pages/Bits.svelte";
import ResubsT1 from "./pages/ResubsT1.svelte";
import Gifts from "./pages/Gifts.svelte";
import ResubsT3 from "./pages/ResubsT3.svelte";
import ResubsT2 from "./pages/ResubsT2.svelte";
import ChatMessages from "./pages/ChatMessages.svelte";

export default [
    {
        title: "Settings",
        path: "/",
        component: Settings,
    },
    {
        title: "Devices",
        path: "/devices",
        component: Devices,
    },
    {
        title: "Settings",
        path: "/settings",
        component: Settings,
    },

    {
        title: "Donations",
        path:"/donations",
        component: Donations,
    },
    {
        title: "Bits",
        path:"/bits",
        component: Bits,
    },
    {
        title: "Gift subs",
        path:"/gifts",
        component: Gifts,
    },
    {
        title: "Resubs Tier 1",
        path:"/resubs1",
        component: ResubsT1,
    },
    {
        title: "Resubs Tier 2",
        path:"/resubs2",
        component: ResubsT2,
    },
    {
        title: "Resubs Tier 3",
        path:"/resubs3",
        component: ResubsT3,
    },
    {
        title: "Chat Messages",
        path: "/chat",
        component: ChatMessages,
    }
];