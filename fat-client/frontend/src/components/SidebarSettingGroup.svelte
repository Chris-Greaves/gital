<script>
    import AdjustableSliders from "../icons/AdjustableSliders.svelte";

    let { name } = $props();
</script>

<li>
    <a
        href="/"
        class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
    >
        <AdjustableSliders style="text-gray-800 dark:text-gray-400" />
        <span class="ms-3">{name}</span>
    </a>
</li>
