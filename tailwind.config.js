module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      colors: {
        // 主色调：深灰/近黑色作为主色
        primary: "#1A1A1A",
        // 次级色：中灰色
        secondary: "#4A4A4A",
        // 强调色：浅灰色
        accent: "#8A8A8A",
        // 中性色系列：从纯白到纯黑的渐变
        neutral: {
          100: "#F5F5F5", // 最浅灰
          200: "#E5E5E5",
          300: "#D4D4D4",
          400: "#A3A3A3",
          500: "#737373",
          600: "#525252",
          700: "#404040",
          800: "#262626",
          900: "#171717", // 近黑色
          950: "#0A0A0A", // 接近纯黑
        },
      },
      fontFamily: {
        sans: ["Inter", "system-ui", "sans-serif"],
        mono: ["JetBrains Mono", "monospace"],
      },
      boxShadow: {
        // 现代简约阴影
        soft: "0 2px 10px rgba(0, 0, 0, 0.05)",
        medium: "0 4px 15px rgba(0, 0, 0, 0.1)",
        hard: "0 8px 30px rgba(0, 0, 0, 0.15)",
      },
    },
  },
  plugins: [],
};
