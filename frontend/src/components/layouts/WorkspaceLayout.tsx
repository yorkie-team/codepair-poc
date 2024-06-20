import * as React from "react";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import { Outlet } from "react-router-dom";
import WorkspaceHeader from "../headers/WorkspaceHeader";
import WorkspaceDrawer from "../drawers/WorkspaceDrawer";

export const DRAWER_WIDTH = 282;

const Main = styled("main", { shouldForwardProp: (prop) => prop !== "open" })<{
	open?: boolean;
}>(({ theme, open }) => ({
	flexGrow: 1,
	padding: theme.spacing(3),
	transition: theme.transitions.create("margin", {
		easing: theme.transitions.easing.sharp,
		duration: theme.transitions.duration.leavingScreen,
	}),
	marginLeft: `-${DRAWER_WIDTH}px`,
	...(open && {
		transition: theme.transitions.create("margin", {
			easing: theme.transitions.easing.easeOut,
			duration: theme.transitions.duration.enteringScreen,
		}),
		marginLeft: 0,
	}),
}));

export const WorkspaceDrawerHeader = styled("div")(({ theme }) => ({
	display: "flex",
	alignItems: "center",
	padding: theme.spacing(0, 1),
	// necessary for content to be below app bar
	...theme.mixins.toolbar,
	justifyContent: "flex-end",
}));

function WorkspaceLayout() {
	const [open, setOpen] = React.useState(false);

	const handleDrawerOpen = () => {
		setOpen((prev) => !prev);
	};

	return (
		<Box sx={{ display: "flex" }}>
			<WorkspaceHeader open={open} onDrawerOpen={handleDrawerOpen} />
			<WorkspaceDrawer open={open} />
			<Main open={open}>
				<WorkspaceDrawerHeader />
				<Outlet />
			</Main>
		</Box>
	);
}

export default WorkspaceLayout;
