import { ReturnCodeLookupRequest } from "../proto/rc";
import { ReturnCodeLookupResult } from "../proto/rc";
import { ReturnCodeLookupResponse } from "../proto/rc";

import "./styles.css";

// Ensure this function is globally accessible if using `onclick` in HTML
export async function lookupReturnCode() {
  // Get the input element and cast it to HTMLInputElement
  const inputElement = document.getElementById(
    "rc-query-input"
  ) as HTMLInputElement;

  const resultsElement = document.getElementById(
    "rc-results"
  ) as HTMLDivElement;

  const inputValue: string = inputElement.value;

  // All TPM error codes are 3-digit hexadecimal numbers.
  if (inputValue.length != 3) {
    resultsElement.textContent = "";
    return;
  }

  // Define the data payload
  const req: ReturnCodeLookupRequest = {
    query: inputValue,
  };

  try {
    const response = await fetch("/rc/lookup", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-protobuf",
      },
      body: ReturnCodeLookupRequest.toBinary(req).buffer as ArrayBuffer,
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.bytes();

    const rsp = ReturnCodeLookupResponse.fromBinary(result);

    resultsElement.textContent = "";

    rsp.result.forEach((res: ReturnCodeLookupResult) => {
      const newDiv = document.createElement("div");
      newDiv.setAttribute("class", "rc-result");
      newDiv.setAttribute("box-", "square");
      newDiv.textContent = res.description;
      resultsElement.appendChild(newDiv);
    });
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to post data.");
  }
}

document.addEventListener("DOMContentLoaded", (event) => {
  const queryInput = document.getElementById(
    "rc-query-input"
  ) as HTMLInputElement | null;

  if (queryInput) {
    queryInput.oninput = lookupReturnCode;
  }
});
