hotkeys("ctrl+space", { splitKey: "+" }, function(event, handler) {
  console.debug(handler.key, event.target, event.key, event.ctrlKey);
});

hotkeys("enter", {}, function(event, handler) {
  // Create new paragraph, with content editable and focus it.
  body = document.querySelector("body");
  body.setAttribute("tabindex", "0");

  // Get which cell we are on now
  p = document.createElement("p");
  body.appendChild(p);
  p.setAttribute("contenteditable", true);
  p.setAttribute("tabindex", "0");

  // Remove, add contenteditable with esc if focused
  p.onkeydown = function(evt) {
    var isEscape = false;
    var isEnter = false;
    // Get correct escape key
    if ("key" in evt) {
      isEscape = evt.key === "Escape" || evt.key === "Esc";
      isEnter = evt.key === "Enter";
    } else {
      isEscape = evt.keyCode === 27;
      isEnter = evt.keyCode === 13;
    }
    // Remove editable
    if (isEscape) {
      p.setAttribute("contenteditable", false);
      body.focus();
    }
    // Add editable
    if (isEnter) {
      p.setAttribute("contenteditable", true);
    }
  };

  p.focus();
});
