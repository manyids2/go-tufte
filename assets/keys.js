hotkeys("ctrl+space", { splitKey: "+" }, function(event, handler) {
  console.debug(handler.key, event.target, event.key, event.ctrlKey);
});
