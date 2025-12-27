import { lookupReturnCode } from "./rc";
import { lookupCommands } from "./cmd";

import "./styles.css";

document.addEventListener("DOMContentLoaded", (event) => {
  const queryInput = document.getElementById(
    "rc-query-input"
  ) as HTMLInputElement | null;

  if (queryInput) {
    queryInput.oninput = lookupReturnCode;
  }

  // TODO: This might be made less janky.
  const cmdResultsElement = document.getElementById(
    "cmd-list-results"
  ) as HTMLDivElement;

  if (cmdResultsElement) {
    lookupCommands();
  }
});
