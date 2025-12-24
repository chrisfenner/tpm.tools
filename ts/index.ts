import { ReturnCodeLookupRequest } from "../proto/rc";
import { ReturnCodeLookupResult } from "../proto/rc";
import { ReturnCodeLookupResponse } from "../proto/rc";

// Ensure this function is globally accessible if using `onclick` in HTML
export async function lookupReturnCode() {
  // Get the input element and cast it to HTMLInputElement
  const inputElement = document.getElementById(
    "rc-query-input"
  ) as HTMLInputElement;
  const inputValue: string = inputElement.value;

  // Define the data payload
  const req: ReturnCodeLookupRequest = {
    query: inputValue,
  };

  let reqBytes: ArrayBuffer = ReturnCodeLookupRequest.toBinary(req)
    .buffer as ArrayBuffer;

  try {
    const response = await fetch("/rc/lookup", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-protobuf",
      },
      body: reqBytes,
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.bytes();

    const rsp = ReturnCodeLookupResponse.fromBinary(result);

    const resultsElement = document.getElementById(
      "rc-results"
    ) as HTMLDivElement;

    resultsElement.textContent = "";

    rsp.result.forEach((res: ReturnCodeLookupResult) => {
      const newDiv = document.createElement("div");
      newDiv.setAttribute("class", "rc-result");

      const newDivCodeName = document.createElement("span");
      newDivCodeName.textContent = "Name: " + res.name;
      const newDivCodeValue = document.createElement("span");
      newDivCodeValue.textContent = "Value: " + res.value;

      newDiv.appendChild(newDivCodeName);
      newDiv.appendChild(newDivCodeValue);
      resultsElement.appendChild(newDiv);
    });
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to post data.");
  }
}

document.addEventListener("DOMContentLoaded", (event) => {
  const queryButton = document.getElementById(
    "rc-query-button"
  ) as HTMLButtonElement | null;

  if (queryButton) {
    queryButton.onclick = lookupReturnCode;
  }
});
