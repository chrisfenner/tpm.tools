import { ReturnCodeLookupRequest, ReturnCodeLookupResponse } from "../proto/rc";
import {
  GetAllCommandNamesRequest,
  GetAllCommandNamesResponse,
  GetCommandDescriptionRequest,
  GetCommandDescriptionResponse,
} from "../proto/cmd";

// Internal helper function for the other functions in this file.
// Send some bytes and get some bytes back.
async function sendReceiveBytes(
  url: string,
  req: Uint8Array
): Promise<Uint8Array> {
  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-protobuf",
    },
    body: req.buffer as ArrayBuffer,
  });

  if (!response.ok) {
    throw new Error(`HTTP status: ${response.status}`);
  }

  return await response.bytes();
}

// The /rc/lookup call.
export async function rcLookup(
  req: ReturnCodeLookupRequest
): Promise<ReturnCodeLookupResponse> {
  return ReturnCodeLookupResponse.fromBinary(
    await sendReceiveBytes("/rc/lookup", ReturnCodeLookupRequest.toBinary(req))
  );
}

// The /cmd/list call.
export async function cmdList(
  req: GetAllCommandNamesRequest
): Promise<GetAllCommandNamesResponse> {
  return GetAllCommandNamesResponse.fromBinary(
    await sendReceiveBytes("/cmd/list", GetAllCommandNamesRequest.toBinary(req))
  );
}

// The /cmd/lookup call.
export async function cmdLookup(
  req: GetCommandDescriptionRequest
): Promise<GetCommandDescriptionResponse> {
  return GetCommandDescriptionResponse.fromBinary(
    await sendReceiveBytes(
      "/cmd/lookup",
      GetCommandDescriptionRequest.toBinary(req)
    )
  );
}
