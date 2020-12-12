"use strict";

document.addEventListener("DOMContentLoaded", setup, false);

function setup() {
    document.querySelector("#submit").onclick = submit;
}

function submit() {
    const post = {
        title: document.querySelector("#title").value,
        id: parseInt(document.querySelector("#id").value),
        markdown: document.querySelector("textarea").value
    };
    fetch("/api/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(post),
    })
}