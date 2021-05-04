import "./index.scss";

import * as React from "react";
import * as ReactDOM from "react-dom";
import GetHighestPrimeFormWidget from "./highestPrime.component";

ReactDOM.render(
    <GetHighestPrimeFormWidget />,
    document.getElementById('root'),
)

// Check that service workers are supported
if ('serviceWorker' in navigator) {
    // Use the window load event to keep the page load performant
    window.addEventListener('load', () => {
        navigator.serviceWorker.register(`${window.location.origin}/public/sw.js`);
    });
}