package good

import (
	trace1 "go.opentelemetry.io/otel/trace"

	goodmgr "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func trace(span trace1.Span, in *npool.GoodReq, index int) trace1.Span {
	return span
}

func Trace(span trace1.Span, in *npool.GoodReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *goodmgr.Conds) trace1.Span {
	return span
}

func TraceMany(span trace1.Span, infos []*npool.GoodReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
