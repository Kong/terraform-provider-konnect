package hooks

func initHooks(h *Hooks) {
    myHook := &FeaturesHook{}
    h.registerBeforeRequestHook(myHook)
}
