import { lookupReturnCode } from "./rc";

import "./styles.css";

document.addEventListener("DOMContentLoaded", (event) => {
  const queryInput = document.getElementById(
    "rc-query-input"
  ) as HTMLInputElement | null;

  if (queryInput) {
    queryInput.oninput = lookupReturnCode;
  }
});
