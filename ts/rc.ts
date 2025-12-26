import { ReturnCodeLookupResult } from "../proto/rc";
import { rcLookup } from "./rest";

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

  try {
    const rsp = await rcLookup({
      query: inputValue,
    });

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
