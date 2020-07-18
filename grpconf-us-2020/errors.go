func CustomErr(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	st := status.Convert(err)
	w.WriteHeader(runtime.HTTPStatusFromCode(st.Code()))
	w.Write([]byte(st.Message()))
}
