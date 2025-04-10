<script>
    import SidebarRepo from "./components/SidebarRepo.svelte";
    import SidebarSettingGroup from "./components/SidebarSettingGroup.svelte";
</script>

<aside
    id="sidebar"
    class="w-64 py-4 px-3 min-h-svh overflow-y-auto bg-gray-50 rounded-sm dark:bg-gray-800"
    aria-label="Sidebar"
>
    <div class="">
        <h2 class="border-b-2 border-b-cyan-900 my-1">Settings</h2>
        <ul class="space-y-2">
            <SidebarSettingGroup name="General" />
        </ul>
    </div>
    <div class="overflow-y-auto bg-gray-50 rounded-sm dark:bg-gray-800">
        <h2 class="border-b-2 border-b-cyan-900 my-1">Repos</h2>
        <ul class="space-y-2">
            <!-- <li>
                <a
                    href=""
                    class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
                    ><svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        class="shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        aria-label="grid solid"
                        viewBox="0 0 24 24"
                        ><path
                            fill-rule="evenodd"
                            d="M4.857 3A1.857 1.857 0 0 0 3 4.857v4.286C3 10.169 3.831 11 4.857 11h4.286A1.857 1.857 0 0 0 11 9.143V4.857A1.857 1.857 0 0 0 9.143 3H4.857Zm10 0A1.857 1.857 0 0 0 13 4.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 21 9.143V4.857A1.857 1.857 0 0 0 19.143 3h-4.286Zm-10 10A1.857 1.857 0 0 0 3 14.857v4.286C3 20.169 3.831 21 4.857 21h4.286A1.857 1.857 0 0 0 11 19.143v-4.286A1.857 1.857 0 0 0 9.143 13H4.857Zm10 0A1.857 1.857 0 0 0 13 14.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 21 19.143v-4.286A1.857 1.857 0 0 0 19.143 13h-4.286Z"
                            clip-rule="evenodd"
                        ></path></svg
                    >
                    <span class="flex-1 ms-3 whitespace-nowrap">Kanban</span>
                    <span
                        class="inline-flex justify-center items-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-200 rounded-full dark:bg-gray-700 dark:text-gray-300"
                        >Pro</span
                    ></a
                >
            </li> -->

            <SidebarRepo pathName="/path/to/repo" />
        </ul>
    </div>
</aside>

<style>
    #sidebar h2 {
        font-size: large;
    }
</style>
