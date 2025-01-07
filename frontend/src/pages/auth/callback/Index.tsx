import { Box } from "@mui/material";
import { useEffect } from "react";
import { useDispatch } from "react-redux";
import { useNavigate, useSearchParams } from "react-router-dom";
import { setAccessToken, setRefreshToken } from "../../../store/authSlice";

function CallbackIndex() {
	const dispatch = useDispatch();
	const navigate = useNavigate();
	const [searchParams] = useSearchParams();

	useEffect(() => {
		const accessToken = searchParams.get("accessToken");
		const refreshToken = searchParams.get("refreshToken");

		if (!accessToken || !refreshToken) {
			navigate("/");
			return;
		}

		dispatch(setAccessToken(accessToken));
		dispatch(setRefreshToken(refreshToken));

		const redirect = sessionStorage.getItem("redirect") || "/";
		sessionStorage.removeItem("redirect");
		navigate(redirect);
	}, [dispatch, navigate, searchParams]);

	return <Box></Box>;
}

export default CallbackIndex;
