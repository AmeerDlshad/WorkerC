package cfruntimecontext

import (
	"context"
	"syscall/js"

	"github.com/AmeerDlshad/WorkerC/internal/runtimecontext"
)

/**
 *  The type definition of RuntimeContext for Cloudflare Worker expects:
 *  ```ts
 *  type RuntimeContext {
 *    env: Env;
 *    ctx: ExecutionContext;
 *  }
 *  ```
 * This type is based on the type definition of ExportedHandlerFetchHandler.
 * - see: https://github.com/cloudflare/workers-types/blob/c8d9533caa4415c2156d2cf1daca75289d01ae70/index.d.ts#LL564
 */

// GetRuntimeContextEnv gets object which holds environment variables bound to Cloudflare worker.
// - see: https://github.com/cloudflare/workers-types/blob/c8d9533caa4415c2156d2cf1daca75289d01ae70/index.d.ts#L566
func GetRuntimeContextEnv(ctx context.Context) js.Value {
	runtimeCtxValue := runtimecontext.MustExtract(ctx)
	return runtimeCtxValue.Get("env")
}

// GetExecutionContext gets ExecutionContext object from context.
// - see: https://github.com/cloudflare/workers-types/blob/c8d9533caa4415c2156d2cf1daca75289d01ae70/index.d.ts#L567
// - see also: https://github.com/cloudflare/workers-types/blob/c8d9533caa4415c2156d2cf1daca75289d01ae70/index.d.ts#L554
func GetExecutionContext(ctx context.Context) js.Value {
	runtimeCtxValue := runtimecontext.MustExtract(ctx)
	return runtimeCtxValue.Get("ctx")
}
