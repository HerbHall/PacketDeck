import { Typography, Box } from "@mui/material";

export function App() {
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" gutterBottom>
        PacketDeck
      </Typography>
      <Typography variant="body1" color="text.secondary">
        Container network topology visualization
      </Typography>
    </Box>
  );
}
