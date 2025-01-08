import KeyboardDoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
import KeyboardDoubleArrowRightIcon from "@mui/icons-material/KeyboardDoubleArrowRight";
import PeopleIcon from "@mui/icons-material/People";
import SpaceDashboardIcon from "@mui/icons-material/SpaceDashboard";
import {
	Box,
	Collapse,
	List,
	ListItemButton,
	ListItemIcon,
	ListItemText,
	Paper,
} from "@mui/material";
import { useMemo, useState } from "react";
import { useDispatch } from "react-redux";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { COLLAPSED_DRAWER_WIDTH, DRAWER_WIDTH } from "../../constants/layout";
import { setDrawerOpen } from "../../store/configSlice";
import { WorkspaceDrawerHeader } from "../layouts/WorkspaceLayout";

interface WorkspaceDrawerProps {
	open: boolean;
}

function WorkspaceDrawer(props: WorkspaceDrawerProps) {
	const { open } = props;
	const location = useLocation();
	const params = useParams();
	const navigate = useNavigate();
	const dispatch = useDispatch();
	const [hovered, setHovered] = useState(false);
	const currentPage = useMemo(() => {
		return location.pathname.split("/")[2] ?? "main";
	}, [location.pathname]);
	const menuList = useMemo(() => {
		return [
			{
				title: "Workspace",
				IconComponent: SpaceDashboardIcon,
				selected: currentPage === "main",
				moveTo: `/${params.workspaceSlug}`,
			},
			{
				title: "Members",
				IconComponent: PeopleIcon,
				selected: currentPage === "member",
				moveTo: `/${params.workspaceSlug}/member`,
			},
		];
	}, [currentPage, params.workspaceSlug]);

	const handleDrawerOpen = () => {
		dispatch(setDrawerOpen(!open));
		setHovered(false);
	};

	const handleMouseEnter = () => {
		setHovered(true);
	};

	const handleMouseLeave = () => {
		setHovered(false);
	};

	return (
		<Box>
			<Box sx={{ width: open ? DRAWER_WIDTH : COLLAPSED_DRAWER_WIDTH }} />
			<Paper
				sx={{
					position: "fixed",
					top: 0,
					left: 0,
					height: "100vh",
					zIndex: 10,
				}}
			>
				<Collapse
					orientation="horizontal"
					in={hovered || open}
					collapsedSize={COLLAPSED_DRAWER_WIDTH}
				>
					<Box
						sx={{
							width: DRAWER_WIDTH,
							overflow: "hidden",
						}}
					>
						<WorkspaceDrawerHeader />
						<Paper
							square
							elevation={0}
							onMouseEnter={handleMouseEnter}
							onMouseLeave={handleMouseLeave}
							sx={{ backgroundColor: "transparent" }}
						>
							<List>
								<ListItemButton onClick={handleDrawerOpen}>
									<ListItemIcon>
										{open ? (
											<KeyboardDoubleArrowLeftIcon />
										) : (
											<KeyboardDoubleArrowRightIcon />
										)}
									</ListItemIcon>
								</ListItemButton>
								{menuList.map((menu, index) => (
									<ListItemButton
										key={index}
										onClick={() => navigate(menu.moveTo)}
										selected={menu.selected}
									>
										<ListItemIcon>
											<menu.IconComponent
												color={menu.selected ? "primary" : "inherit"}
											/>
										</ListItemIcon>
										<ListItemText primary={menu.title} />
									</ListItemButton>
								))}
							</List>
						</Paper>
					</Box>
				</Collapse>
			</Paper>
		</Box>
	);
}

export default WorkspaceDrawer;
