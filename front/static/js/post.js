"use strict";

document.addEventListener("DOMContentLoaded", setup, false);

function setup() {
    fetch("/api/post" + (window.location.pathname === "/" ? "/0" : window.location.pathname))
        .then(response => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error("Post not found");
            }
        })
        .then(post => {
            document.title = post.title;
            document.querySelector("body").innerHTML = post.html;
        })
        .catch(e => {
            document.title = "404";
            document.querySelector("body").innerHTML = "<h1>404 post not found</h1>";
        });
}
