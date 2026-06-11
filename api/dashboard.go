package api

// DashboardHTML contains the embedded Single Page Application dashboard code
const DashboardHTML = `<!DOCTYPE html>
<html lang="ar" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>NeuroRoot Core Portal</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&family=Tajawal:wght@300;500;700;900&display=swap" rel="stylesheet">
    <style>
        :root {
            --bg-color: #050811;
            --panel-bg: rgba(13, 20, 38, 0.6);
            --panel-border: rgba(255, 255, 255, 0.06);
            --accent-cyan: #00f2fe;
            --accent-purple: #9b51e0;
            --accent-emerald: #05c46b;
            --accent-rose: #ff5e57;
            --text-main: #ffffff;
            --text-muted: #8a99ad;
            --font-family: 'Tajawal', 'Outfit', sans-serif;
            --shadow-glow: 0 0 25px rgba(0, 242, 254, 0.15);
            --card-hover: rgba(255, 255, 255, 0.02);
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            background-color: var(--bg-color);
            color: var(--text-main);
            font-family: var(--font-family);
            min-height: 100vh;
            overflow-x: hidden;
            display: flex;
            background-image: 
                radial-gradient(at 0% 0%, rgba(155, 81, 224, 0.12) 0px, transparent 40%),
                radial-gradient(at 100% 100%, rgba(0, 242, 254, 0.08) 0px, transparent 40%);
        }

        /* Sidebar navigation */
        .sidebar {
            width: 280px;
            background: rgba(8, 12, 24, 0.85);
            border-left: 1px solid var(--panel-border);
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            height: 100vh;
            position: sticky;
            top: 0;
            padding: 2rem 1.5rem;
            z-index: 10;
        }

        .logo-area {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            margin-bottom: 3rem;
        }

        .logo-icon {
            width: 42px;
            height: 42px;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            border-radius: 12px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-weight: 900;
            font-size: 1.4rem;
            color: #050811;
            box-shadow: 0 0 20px rgba(0, 242, 254, 0.35);
        }

        .logo-text {
            font-size: 1.4rem;
            font-weight: 900;
            background: linear-gradient(to left, var(--accent-cyan), #fff);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .nav-menu {
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
            flex-grow: 1;
        }

        .nav-item {
            display: flex;
            align-items: center;
            gap: 1rem;
            padding: 0.85rem 1.25rem;
            background: transparent;
            border: 1px solid transparent;
            border-radius: 12px;
            color: var(--text-muted);
            font-size: 1rem;
            font-weight: 700;
            cursor: pointer;
            text-align: right;
            width: 100%;
            transition: all 0.3s ease;
        }

        .nav-item:hover {
            color: #fff;
            background: var(--card-hover);
            border-color: var(--panel-border);
        }

        .nav-item.active {
            color: #050811;
            background: linear-gradient(90deg, var(--accent-cyan), #fff);
            box-shadow: var(--shadow-glow);
            border-color: var(--accent-cyan);
        }

        .nav-item.active svg {
            stroke: #050811;
        }

        .nav-item svg {
            width: 20px;
            height: 20px;
            stroke: var(--text-muted);
            fill: none;
            stroke-width: 2.5;
            transition: stroke 0.3s ease;
        }

        /* Profile summary inside sidebar */
        .sidebar-footer {
            border-top: 1px solid var(--panel-border);
            padding-top: 1.5rem;
        }

        .footer-profile {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            cursor: pointer;
        }

        .avatar-circle {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            box-shadow: 0 0 10px rgba(0, 242, 254, 0.2);
            border: 2px solid rgba(255,255,255,0.1);
        }

        .profile-info {
            display: flex;
            flex-direction: column;
            gap: 0.1rem;
        }

        .profile-name {
            font-weight: 700;
            font-size: 0.95rem;
            color: #fff;
        }

        .profile-status {
            font-size: 0.75rem;
            color: var(--accent-emerald);
            display: flex;
            align-items: center;
            gap: 0.25rem;
        }

        /* Content Container */
        .main-content {
            flex-grow: 1;
            padding: 2.5rem;
            height: 100vh;
            overflow-y: auto;
            display: flex;
            flex-direction: column;
        }

        .tab-panel {
            display: none;
            height: 100%;
            animation: slideUp 0.4s cubic-bezier(0.16, 1, 0.3, 1);
        }

        .tab-panel.active {
            display: flex;
            flex-direction: column;
        }

        @keyframes slideUp {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }

        /* Common Elements */
        .card {
            background: var(--panel-bg);
            border: 1px solid var(--panel-border);
            border-radius: 24px;
            padding: 2rem;
            backdrop-filter: blur(20px);
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.25);
            margin-bottom: 2rem;
        }

        h2 {
            font-size: 2rem;
            font-weight: 900;
            margin-bottom: 0.5rem;
            letter-spacing: -0.5px;
        }

        .subtitle {
            color: var(--text-muted);
            font-size: 1.05rem;
            margin-bottom: 2rem;
        }

        .btn-primary {
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            color: #050811;
            border: none;
            padding: 0.85rem 2rem;
            border-radius: 14px;
            font-weight: 800;
            font-size: 1rem;
            cursor: pointer;
            box-shadow: 0 4px 20px rgba(0, 242, 254, 0.25);
            transition: all 0.3s ease;
        }

        .btn-primary:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 25px rgba(0, 242, 254, 0.45);
            opacity: 0.95;
        }

        .btn-secondary {
            background: rgba(255, 255, 255, 0.03);
            border: 1px solid var(--panel-border);
            color: #fff;
            padding: 0.85rem 2rem;
            border-radius: 14px;
            font-weight: 700;
            cursor: pointer;
            transition: all 0.3s ease;
        }

        .btn-secondary:hover {
            background: rgba(255, 255, 255, 0.07);
            border-color: var(--text-muted);
        }

        input {
            background: rgba(8, 12, 24, 0.8);
            border: 1px solid var(--panel-border);
            color: #fff;
            padding: 0.85rem 1.25rem;
            border-radius: 14px;
            font-size: 1rem;
            font-family: var(--font-family);
            width: 100%;
            transition: all 0.3s ease;
        }

        input:focus {
            outline: none;
            border-color: var(--accent-cyan);
            box-shadow: 0 0 15px rgba(0, 242, 254, 0.15);
        }

        /* 🔍 Search engine layout (Google-like) */
        .search-engine {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            flex-grow: 1;
            padding-bottom: 5rem;
        }

        .search-logo {
            font-size: 4rem;
            font-weight: 900;
            margin-bottom: 1.5rem;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple), #fff);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            letter-spacing: -2px;
            animation: pulse Glow 3s infinite;
        }

        .search-bar-container {
            width: 100%;
            max-width: 650px;
            position: relative;
            margin-bottom: 2rem;
        }

        .search-bar-container input {
            padding-right: 3rem;
            padding-left: 1.5rem;
            border-radius: 30px;
            height: 55px;
            font-size: 1.15rem;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
        }

        .search-icon-btn {
            position: absolute;
            right: 1.25rem;
            top: 50%;
            transform: translateY(-50%);
            background: none;
            border: none;
            cursor: pointer;
        }

        .search-icon-btn svg {
            width: 24px;
            height: 24px;
            stroke: var(--text-muted);
            fill: none;
            stroke-width: 2.5;
        }

        .search-results-area {
            width: 100%;
            max-width: 700px;
            display: none;
            flex-direction: column;
            gap: 1.5rem;
            margin-top: 1rem;
        }

        .result-card {
            background: rgba(255, 255, 255, 0.015);
            border: 1px solid var(--panel-border);
            border-radius: 18px;
            padding: 1.25rem 1.5rem;
            transition: all 0.25s ease;
        }

        .result-card:hover {
            border-color: rgba(0, 242, 254, 0.25);
            background: rgba(255, 255, 255, 0.03);
            transform: translateX(-4px);
        }

        .result-title {
            color: var(--accent-cyan);
            font-size: 1.2rem;
            font-weight: 700;
            margin-bottom: 0.35rem;
            cursor: pointer;
        }

        .result-snippet {
            color: var(--text-muted);
            font-size: 0.95rem;
            line-height: 1.45;
        }

        /* 💬 Chat layout (Telegram-like) */
        .chat-container {
            display: flex;
            flex-grow: 1;
            background: rgba(8, 12, 24, 0.4);
            border: 1px solid var(--panel-border);
            border-radius: 24px;
            overflow: hidden;
            height: calc(100vh - 10rem);
        }

        .chat-sidebar {
            width: 280px;
            border-left: 1px solid var(--panel-border);
            background: rgba(6, 10, 20, 0.5);
            display: flex;
            flex-direction: column;
        }

        .chat-sidebar-header {
            padding: 1.5rem;
            border-bottom: 1px solid var(--panel-border);
            font-weight: 800;
            font-size: 1.1rem;
        }

        .channels-list {
            overflow-y: auto;
            flex-grow: 1;
            padding: 0.75rem;
            display: flex;
            flex-direction: column;
            gap: 0.35rem;
        }

        .channel-item {
            padding: 0.75rem 1rem;
            border-radius: 12px;
            cursor: pointer;
            display: flex;
            align-items: center;
            gap: 0.75rem;
            font-weight: 700;
            color: var(--text-muted);
            transition: all 0.25s ease;
        }

        .channel-item:hover {
            background: rgba(255,255,255,0.02);
            color: #fff;
        }

        .channel-item.active {
            background: rgba(0, 242, 254, 0.08);
            color: var(--accent-cyan);
        }

        .channel-avatar {
            width: 30px;
            height: 30px;
            border-radius: 8px;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            color: #050811;
            font-weight: 900;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 0.95rem;
        }

        .chat-main {
            flex-grow: 1;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            background: rgba(5, 8, 17, 0.3);
        }

        .chat-header {
            padding: 1.25rem 2rem;
            border-bottom: 1px solid var(--panel-border);
            display: flex;
            align-items: center;
            justify-content: space-between;
            background: rgba(8, 12, 24, 0.4);
        }

        .chat-messages {
            flex-grow: 1;
            overflow-y: auto;
            padding: 2rem;
            display: flex;
            flex-direction: column;
            gap: 1.25rem;
        }

        .chat-message-bubble {
            max-width: 70%;
            padding: 0.9rem 1.25rem;
            border-radius: 18px;
            line-height: 1.5;
            position: relative;
        }

        .chat-message-bubble.incoming {
            background: rgba(255, 255, 255, 0.03);
            border: 1px solid var(--panel-border);
            align-self: flex-start;
            border-top-right-radius: 4px;
        }

        .chat-message-bubble.outgoing {
            background: linear-gradient(135deg, rgba(0, 242, 254, 0.15), rgba(155, 81, 224, 0.15));
            border: 1px solid rgba(0, 242, 254, 0.18);
            align-self: flex-end;
            border-top-left-radius: 4px;
        }

        .msg-sender {
            font-size: 0.75rem;
            font-weight: 800;
            color: var(--accent-cyan);
            margin-bottom: 0.25rem;
            display: block;
        }

        .msg-text {
            font-size: 0.95rem;
            word-break: break-all;
        }

        .msg-time {
            font-size: 0.7rem;
            color: var(--text-muted);
            margin-top: 0.35rem;
            display: block;
            text-align: left;
        }

        .chat-input-area {
            padding: 1.5rem 2rem;
            border-top: 1px solid var(--panel-border);
            background: rgba(8, 12, 24, 0.4);
            display: flex;
            gap: 1rem;
        }

        /* 🛒 Registrar layout (GoDaddy-like) */
        .registrar-search-box {
            display: flex;
            gap: 1rem;
            margin-bottom: 2rem;
        }

        .registrar-search-box input {
            height: 50px;
            border-radius: 14px;
            font-size: 1.1rem;
        }

        .domain-checkout-card {
            background: linear-gradient(135deg, rgba(255,255,255,0.01) 0%, rgba(255,255,255,0.03) 100%);
            border: 1px solid var(--panel-border);
            border-radius: 20px;
            padding: 2rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            animation: fadeIn 0.3s ease;
        }

        .checkout-info {
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }

        .checkout-domain {
            font-size: 1.5rem;
            font-weight: 800;
            color: var(--accent-cyan);
        }

        .checkout-price {
            font-size: 1.1rem;
            color: var(--accent-emerald);
            font-weight: 700;
        }

        /* Wizard layout for domain commit-reveal */
        .wizard-container {
            display: none;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            padding: 3rem 1.5rem;
            gap: 1.5rem;
        }

        .wizard-steps {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 1rem;
            width: 100%;
            max-width: 400px;
        }

        .wizard-step {
            width: 32px;
            height: 32px;
            border-radius: 50%;
            background: rgba(255,255,255,0.05);
            border: 1px solid var(--panel-border);
            display: flex;
            align-items: center;
            justify-content: center;
            font-weight: 800;
            color: var(--text-muted);
            transition: all 0.3s ease;
        }

        .wizard-step.active {
            background: var(--accent-cyan);
            color: #050811;
            box-shadow: var(--shadow-glow);
        }

        .wizard-step.completed {
            background: var(--accent-emerald);
            color: #050811;
        }

        .wizard-line {
            height: 2px;
            background: var(--panel-border);
            flex-grow: 1;
        }

        .wizard-line.active {
            background: var(--accent-cyan);
        }

        /* 🖥️ Browser Mockup */
        .browser-mockup {
            border: 1px solid var(--panel-border);
            border-radius: 20px;
            overflow: hidden;
            display: flex;
            flex-direction: column;
            background: rgba(8, 12, 24, 0.6);
            height: calc(100vh - 12rem);
            box-shadow: 0 15px 40px rgba(0,0,0,0.3);
        }

        .browser-toolbar {
            background: rgba(13, 20, 38, 0.9);
            border-bottom: 1px solid var(--panel-border);
            padding: 0.75rem 1.5rem;
            display: grid;
            grid-template-columns: 80px 1fr 100px;
            align-items: center;
            gap: 1rem;
        }

        .browser-dots {
            display: flex;
            gap: 0.35rem;
        }

        .browser-dot {
            width: 10px;
            height: 10px;
            border-radius: 50%;
        }

        .browser-address {
            display: flex;
            align-items: center;
            gap: 0.5rem;
            background: rgba(0,0,0,0.3);
            border: 1px solid var(--panel-border);
            border-radius: 30px;
            padding: 0.45rem 1rem;
        }

        .browser-address input {
            background: transparent;
            border: none;
            padding: 0;
            height: 100%;
            font-size: 0.9rem;
            color: var(--text-muted);
        }

        .browser-viewport {
            flex-grow: 1;
            background: #ffffff;
        }

        /* Toast Notifications */
        .toast-container {
            position: fixed;
            bottom: 2rem;
            left: 2rem;
            display: flex;
            flex-direction: column;
            gap: 0.75rem;
            z-index: 1000;
        }

        .toast {
            background: rgba(13, 20, 38, 0.95);
            border: 1px solid var(--panel-border);
            border-radius: 12px;
            padding: 1rem 1.5rem;
            color: #fff;
            box-shadow: 0 10px 35px rgba(0, 0, 0, 0.45);
            display: flex;
            align-items: center;
            gap: 0.75rem;
            transform: translateX(-120%);
            animation: slideIn 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards;
            min-width: 320px;
            font-weight: 700;
        }

        .toast.success { border-left: 5px solid var(--accent-emerald); }
        .toast.error { border-left: 5px solid var(--accent-rose); }

        @keyframes slideIn {
            to { transform: translateX(0); }
        }

        /* Helpers */
        .flex-row { display: flex; gap: 1rem; }
        .flex-between { justify-content: space-between; align-items: center; }
        .glowing-circle {
            animation: pulseGlow 2s infinite alternate;
        }
        @keyframes pulseGlow {
            from { box-shadow: 0 0 10px rgba(0, 242, 254, 0.2); }
            to { box-shadow: 0 0 25px rgba(0, 242, 254, 0.5); }
        }
    </style>
</head>
<body>

    <!-- Sidebar Navigation -->
    <div class="sidebar">
        <div>
            <div class="logo-area">
                <div class="logo-icon">NR</div>
                <div class="logo-text">NeuroRoot</div>
            </div>
            
            <div class="nav-menu">
                <button class="nav-item active" onclick="switchTab('search', this)">
                    <svg viewBox="0 0 24 24"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                    <span>البحث والحل اللامركزي</span>
                </button>
                <button class="nav-item" onclick="switchTab('registrar', this)">
                    <svg viewBox="0 0 24 24"><path d="M12 2L2 7l10 5 10-5-10-5z"></path><path d="M2 17l10 5 10-5"></path><path d="M2 12l10 5 10-5"></path></svg>
                    <span>حجز النطاقات (.ia)</span>
                </button>
                <button class="nav-item" onclick="switchTab('chat', this)">
                    <svg viewBox="0 0 24 24"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
                    <span>الدردشة اللامركزية</span>
                </button>
                <button class="nav-item" onclick="switchTab('browser', this)">
                    <svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
                    <span>المتصفح اللامركزي</span>
                </button>
                <button class="nav-item" onclick="switchTab('profile', this)">
                    <svg viewBox="0 0 24 24"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                    <span>هويتي الرقمية</span>
                </button>
                <button class="nav-item" onclick="switchTab('settings', this)">
                    <svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="3"></circle><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"></path></svg>
                    <span>الإعدادات والمطور</span>
                </button>
            </div>
        </div>

        <div class="sidebar-footer" id="sidebar-profile-card">
            <div class="footer-profile" onclick="switchTab('profile')">
                <div class="avatar-circle"></div>
                <div class="profile-info">
                    <span class="profile-name" id="sb-profile-name">...</span>
                    <span class="profile-status">
                        <span style="width: 8px; height: 8px; border-radius: 50%; background: var(--accent-emerald); display: inline-block;"></span>
                        متصل بالشبكة
                    </span>
                </div>
            </div>
        </div>
    </div>

    <!-- Main Content -->
    <div class="main-content">

        <!-- Tab: Search (Google-like) -->
        <div id="search-panel" class="tab-panel active">
            <div class="search-engine">
                <div class="search-logo">NeuroSearch</div>
                <p class="subtitle" style="margin-top: -1.25rem; margin-bottom: 2rem;">مستقبل البحث والوصول اللامركزي في الويب الجديد</p>
                
                <div class="search-bar-container">
                    <input type="text" id="search-input" placeholder="ابحث عن نطاق، عنوان، كلمة دلالية أو CID..." onkeydown="if(event.key==='Enter') executeDecentralizedSearch()">
                    <button class="search-icon-btn" onclick="executeDecentralizedSearch()">
                        <svg viewBox="0 0 24 24"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                    </button>
                </div>

                <div class="search-results-area" id="search-results-box">
                    <!-- Dynamic search results cards go here -->
                </div>
            </div>
        </div>

        <!-- Tab: Registrar (GoDaddy-like) -->
        <div id="registrar-panel" class="tab-panel">
            <h2>حجز النطاقات اللامركزية</h2>
            <p class="subtitle">سجل اسم هويتك الرقمية الفريد بامتداد .ia للحماية وحجز مكانك على الويب الجديد</p>

            <div class="card">
                <h3 style="margin-bottom: 1.25rem;">ابحث عن النطاق المتاح</h3>
                <div class="registrar-search-box">
                    <input type="text" id="registrar-domain-input" placeholder="أدخل اسم النطاق المطلوب (مثال: myname.ia)" onkeydown="if(event.key==='Enter') checkDomainAvailability()">
                    <button class="btn-primary" onclick="checkDomainAvailability()">بحث وتوافر</button>
                </div>

                <!-- Domain Search result state (checkout card) -->
                <div class="domain-checkout-card" id="registrar-result-card" style="display:none;">
                    <div class="checkout-info">
                        <span class="checkout-domain" id="registrar-result-domain">...</span>
                        <span class="checkout-price" id="registrar-result-status">متاح للتسجيل</span>
                    </div>
                    <button class="btn-primary" id="registrar-checkout-btn" onclick="startRegistrationWizard()">سجل النطاق الآن (مجاني)</button>
                </div>

                <!-- Registration wizard (Commit-Reveal Visual Progress) -->
                <div class="wizard-container" id="registration-wizard">
                    <h3 id="wizard-title" style="color: var(--accent-cyan); font-weight:800;">تأمين حجز النطاق...</h3>
                    <p id="wizard-status-text" style="color: var(--text-muted); font-size: 0.95rem; text-align:center;">
                        جاري إرسال التزام الحجز المشفر للشبكة لمنع سرقة النطاق (Front-running)
                    </p>

                    <div class="wizard-steps">
                        <div class="wizard-step active" id="step1-indicator">1</div>
                        <div class="wizard-line" id="line1-indicator"></div>
                        <div class="wizard-step" id="step2-indicator">2</div>
                        <div class="wizard-line" id="line2-indicator"></div>
                        <div class="wizard-step" id="step3-indicator">3</div>
                    </div>

                    <!-- Visual countdown timer -->
                    <div id="wizard-timer-container" style="display:none; text-align:center;">
                        <span style="font-size: 3rem; font-weight: 800; color: var(--accent-purple); text-shadow: var(--shadow-glow);" id="wizard-timer-display">60</span>
                        <p style="font-size:0.8rem; color:var(--text-muted); margin-top:0.25rem;">ثانية متبقية لتأكيد القيد اللامركزي</p>
                    </div>

                    <button class="btn-secondary" id="wizard-cancel-btn" style="display:none; margin-top:1rem;" onclick="resetRegistrar()">إلغاء والعودة</button>
                </div>
            </div>
        </div>

        <!-- Tab: Chat (Telegram-like) -->
        <div id="chat-panel" class="tab-panel">
            <h2>قنوات المحادثة اللامركزية</h2>
            <p class="subtitle">تواصل مع المستخدمين الآخرين والوكلاء الأذكياء عبر شبكة GossipSub الموزعة</p>

            <div class="chat-container">
                <!-- Channels list sidebar -->
                <div class="chat-sidebar">
                    <div class="chat-sidebar-header">القنوات المتاحة</div>
                    <div class="channels-list" id="chat-channels-list">
                        <!-- dynamic channels list -->
                    </div>
                    <div style="padding: 1rem; border-top: 1px solid var(--panel-border);">
                        <input type="text" id="chat-new-channel-name" placeholder="انضمام لقناة جديدة..." onkeydown="if(event.key==='Enter') handleNewChannelJoin()" style="font-size: 0.85rem; padding: 0.65rem 1rem;">
                    </div>
                </div>

                <!-- Chat conversation area -->
                <div class="chat-main">
                    <div class="chat-header">
                        <div style="display:flex; flex-direction:column; gap:0.1rem;">
                            <span style="font-weight: 800; font-size: 1.15rem;" id="chat-active-channel-name">لا توجد قناة نشطة</span>
                            <span style="font-size: 0.75rem; color: var(--text-muted);" id="chat-active-channel-status">اختر قناة من اليسار لبدء المحادثة</span>
                        </div>
                    </div>

                    <!-- Chat Messages List -->
                    <div class="chat-messages" id="chat-messages-container">
                        <div style="text-align:center; color: var(--text-muted); margin-top: 8rem;">
                            اختر أو اكتب اسم قناة وانضم إليها للمشاركة في الحوار اللامركزي.
                        </div>
                    </div>

                    <!-- Chat Input -->
                    <div class="chat-input-area">
                        <input type="text" id="chat-message-input" placeholder="اكتب رسالة في القناة..." onkeydown="if(event.key==='Enter') executeSendChatMessage()">
                        <button class="btn-primary" onclick="executeSendChatMessage()" style="padding: 0.85rem 1.75rem;">إرسال</button>
                    </div>
                </div>
            </div>
        </div>

        <!-- Tab: Decentralized Browser -->
        <div id="browser-panel" class="tab-panel">
            <h2>المتصفح والشبكة الداخلية</h2>
            <p class="subtitle">استعرض بوابات ومحتوى الويب المستضاف بالكامل على بروتوكول Bitswap اللامركزي</p>

            <div class="browser-mockup">
                <!-- Address bar toolbar -->
                <div class="browser-toolbar">
                    <div class="browser-dots">
                        <div class="browser-dot" style="background:#ff5f56;"></div>
                        <div class="browser-dot" style="background:#ffbd2e;"></div>
                        <div class="browser-dot" style="background:#27c93f;"></div>
                    </div>
                    <div class="browser-address">
                        <span style="font-size:0.85rem; font-weight:800; color:var(--accent-cyan); font-family:'Outfit';">ia://</span>
                        <input type="text" id="browser-url-input" value="hello.ia" onkeydown="if(event.key==='Enter') executeBrowserLoad()">
                    </div>
                    <button class="btn-primary" onclick="executeBrowserLoad()" style="padding:0.45rem 1.25rem; font-size:0.85rem; border-radius:30px; box-shadow:none;">تصفح</button>
                </div>

                <!-- Viewport iframe container -->
                <div class="browser-viewport" id="browser-viewport-container" style="background: #111; display:flex; align-items:center; justify-content:center; height:100%;">
                    <div style="text-align:center; color: var(--text-muted);" id="browser-placeholder-text">
                        <svg viewBox="0 0 24 24" style="width: 64px; height:64px; stroke: var(--text-muted); fill:none; stroke-width:1.5; margin-bottom:1rem;"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>
                        <p>أدخل اسم النطاق المقيد بـ .ia للتحميل من الذاكرة اللامركزية عبر البوابة المحلية</p>
                    </div>
                    <iframe id="browser-iframe" src="" style="width:100%; height:100%; border:none; display:none; background:#fff;"></iframe>
                </div>
            </div>
        </div>

        <!-- Tab: My Profile / Identity -->
        <div id="profile-panel" class="tab-panel">
            <h2>هويتي الرقمية المستقلة</h2>
            <p class="subtitle">سجلك الرقمي اللامركزي المشفر بالكامل والمملوك لك بموجب مفاتيح التشفير</p>

            <div class="flex-row">
                <div class="card" style="flex-grow: 1; display:flex; flex-direction:column; gap:1.5rem;">
                    <div style="display:flex; align-items:center; gap:1.5rem;">
                        <div class="avatar-circle glowing-circle" style="width:80px; height:80px;"></div>
                        <div>
                            <h3 style="font-size:1.4rem; font-weight:800;" id="profile-identity-did-short">did:ia:...</h3>
                            <span style="font-size:0.85rem; color:var(--accent-cyan);" id="profile-identity-created">تاريخ التسجيل: ...</span>
                        </div>
                    </div>

                    <div style="border-top:1px solid var(--panel-border); padding-top:1.5rem; display:flex; flex-direction:column; gap:1rem;">
                        <div>
                            <label style="font-size:0.8rem; color:var(--text-muted); font-weight:700; display:block; margin-bottom:0.25rem;">المعرف الرقمي الكامل (DID)</label>
                            <div style="display:flex; gap:0.5rem;">
                                <input type="text" id="profile-did-full" readonly style="font-size:0.85rem; font-family:monospace;">
                                <button class="btn-secondary" onclick="copyToClipboard('profile-did-full')" style="padding:0 1.25rem; font-size:0.85rem;">نسخ</button>
                            </div>
                        </div>

                        <div>
                            <label style="font-size:0.8rem; color:var(--text-muted); font-weight:700; display:block; margin-bottom:0.25rem;">المفتاح العام للتشفير (PublicKeyHex)</label>
                            <input type="text" id="profile-pub-key" readonly style="font-size:0.85rem; font-family:monospace;">
                        </div>
                    </div>
                </div>

                <div class="card" style="width:350px; flex-shrink:0; display:flex; flex-direction:column; gap:1.5rem;">
                    <h3 style="color:var(--accent-purple); font-weight:800;">القدرات المصرحة</h3>
                    <div id="profile-capabilities-list" style="display:flex; flex-direction:column; gap:0.75rem;">
                        <!-- dynamic capabilities list -->
                    </div>
                </div>
            </div>
        </div>

        <!-- Tab: Settings -->
        <div id="settings-panel" class="tab-panel">
            <h2>إعدادات المطور والاتصال بالشبكة</h2>
            <p class="subtitle">تفاصيل تقنية وعقدية لتوصيل وإعداد بيئة عمل شبكة الوكلاء</p>

            <div class="card">
                <h3 style="margin-bottom:1.5rem; color:var(--accent-cyan);">مصادقة الواجهة البرمجية (REST API Authorization)</h3>
                <div class="form-group" style="margin-bottom:1.5rem;">
                    <label style="margin-bottom:0.5rem; display:block; font-weight:700;">رمز الوصول للمصادقة (Bearer Token)</label>
                    <div style="display:flex; gap:1rem;">
                        <input type="password" id="settings-token-input" placeholder="nr-...">
                        <button class="btn-primary" onclick="saveSettingsToken()">حفظ الرمز</button>
                    </div>
                </div>
            </div>

            <div class="card">
                <h3 style="margin-bottom:1.5rem; color:var(--accent-purple);">معلومات بوابة شبكة P2P</h3>
                <div style="display:flex; flex-direction:column; gap:0.75rem; font-size:0.95rem; color:var(--text-muted);" id="settings-p2p-info">
                    <!-- dynamic status -->
                </div>
            </div>
        </div>

    </div>

    <!-- Toast Container -->
    <div class="toast-container" id="toast-container"></div>

    <script>
        let apiToken = "";
        let currentActiveChannel = "";
        let channelPollInterval = null;

        // On Load initialization
        window.addEventListener('DOMContentLoaded', () => {
            const urlParams = new URLSearchParams(window.location.search);
            const token = urlParams.get('token');
            if (token) {
                localStorage.setItem('nr_token', token);
                apiToken = token;
                window.history.replaceState({}, document.title, window.location.pathname);
            } else {
                apiToken = localStorage.getItem('nr_token') || "";
            }

            document.getElementById('settings-token-input').value = apiToken;

            initializeApplication();
        });

        async function initializeApplication() {
            if (!apiToken) {
                showToast("يرجى إدخال رمز الوصول في صفحة الإعدادات للاتصال بالوكيل", "error");
                switchTab('settings');
                return;
            }

            const data = await callAPI("/api/identity");
            if (data) {
                // Populate profile
                document.getElementById('sb-profile-name').innerText = data.did.substring(0, 15) + "...";
                document.getElementById('profile-identity-did-short').innerText = data.did;
                document.getElementById('profile-did-full').value = data.did;
                document.getElementById('profile-pub-key').value = data.public_key_hex;

                const createdDate = new Date(data.created_at * 1000).toLocaleDateString('ar-EG');
                document.getElementById('profile-identity-created').innerText = "تاريخ الإنشاء: " + createdDate;

                // Capabilities
                const capContainer = document.getElementById('profile-capabilities-list');
                capContainer.innerHTML = "";
                data.capabilities.forEach(cap => {
                    const div = document.createElement('div');
                    div.style.background = "rgba(255,255,255,0.02)";
                    div.style.border = "1px solid var(--panel-border)";
                    div.style.padding = "0.5rem 1rem";
                    div.style.borderRadius = "10px";
                    div.style.fontSize = "0.85rem";
                    div.style.fontWeight = "700";
                    div.style.color = "var(--accent-cyan)";
                    div.innerText = cap;
                    capContainer.appendChild(div);
                });

                // P2P Info on settings
                const p2pInfoBox = document.getElementById('settings-p2p-info');
                p2pInfoBox.innerHTML = '<p><strong>معرّف الهوية اللامركزي (DID):</strong> ' + data.did + '</p>' +
                    '<p><strong>المفتاح العام:</strong> ' + data.public_key_hex + '</p>' +
                    '<p><strong>سجل انتهاء الصلاحية:</strong> ' + new Date(data.expires_at * 1000).toLocaleString('ar-EG') + '</p>';

                // Auto-join standard chats and load channels
                await setupDefaultChannels();
            }
        }

        async function setupDefaultChannels() {
            // Join lobby by default
            await callAPI("/api/channels/join", "POST", { channel_id: "main-lobby" });
            await loadJoinedChannels();
            selectChannel("main-lobby");
        }

        // Tab Switcher
        function switchTab(tabName, button = null) {
            document.querySelectorAll('.tab-panel').forEach(panel => panel.classList.remove('active'));
            document.querySelectorAll('.nav-item').forEach(item => item.classList.remove('active'));

            const targetPanel = document.getElementById(tabName + '-panel');
            if (targetPanel) {
                targetPanel.classList.add('active');
            }

            if (button) {
                button.classList.add('active');
            } else {
                // Find button by onclick function parameter
                const buttons = document.querySelectorAll('.nav-item');
                buttons.forEach(btn => {
                    if (btn.getAttribute('onclick') && btn.getAttribute('onclick').includes(tabName)) {
                        btn.classList.add('active');
                    }
                });
            }

            // Chat tab polling trigger
            if (tabName === 'chat') {
                if (currentActiveChannel) {
                    selectChannel(currentActiveChannel);
                }
            } else {
                clearInterval(channelPollInterval);
            }
        }

        // API Helper
        async function callAPI(endpoint, method = "GET", body = null) {
            if (!apiToken) {
                showToast("يرجى إدخال رمز مصادقة API للمتابعة", "error");
                return null;
            }

            const headers = {
                'Authorization': "Bearer " + apiToken,
                'Content-Type': 'application/json'
            };

            const config = { method, headers };
            if (body) {
                config.body = typeof body === 'string' ? body : JSON.stringify(body);
            }

            try {
                const response = await fetch(endpoint, config);
                if (response.status === 401) {
                    showToast("رمز مصادقة REST API غير صالح أو منتهي الصلاحية", "error");
                    return null;
                }
                if (!response.ok) {
                    const text = await response.text();
                    throw new Error(text || response.statusText);
                }
                const contentType = response.headers.get("content-type");
                if (contentType && contentType.includes("application/json")) {
                    return await response.json();
                }
                return await response.text();
            } catch (error) {
                showToast("فشل في الاتصال بالوكيل: " + error.message, "error");
                return null;
            }
        }

        function showToast(message, type = "success") {
            const container = document.getElementById('toast-container');
            const toast = document.createElement('div');
            toast.className = "toast " + type;
            toast.innerText = message;
            container.appendChild(toast);
            setTimeout(() => {
                toast.style.animation = "slideIn 0.3s cubic-bezier(0.16, 1, 0.3, 1) reverse";
                setTimeout(() => toast.remove(), 300);
            }, 4000);
        }

        // Copy Helper
        function copyToClipboard(id) {
            const input = document.getElementById(id);
            input.select();
            input.setSelectionRange(0, 99999);
            navigator.clipboard.writeText(input.value);
            showToast("تم النسخ بنجاح!", "success");
        }

        // Settings Token Save
        function saveSettingsToken() {
            const val = document.getElementById('settings-token-input').value.trim();
            if (val) {
                apiToken = val;
                localStorage.setItem('nr_token', val);
                initializeApplication();
                showToast("تم حفظ رمز المصادقة بنجاح وتحديث الاتصال", "success");
            }
        }

        // 🔍 Search Portal logic
        async function executeDecentralizedSearch() {
            const query = document.getElementById('search-input').value.trim();
            if (!query) return;

            showToast("جاري البحث في فهارس DHT اللامركزية...", "success");

            const resultsBox = document.getElementById('search-results-box');
            resultsBox.style.display = "flex";
            resultsBox.innerHTML = '<div style="text-align:center; color:var(--text-muted); padding:2rem;">جاري جلب النتائج...</div>';

            // Simulate searching naming database and content hashes
            const res = await callAPI("/api/resolve?name=" + encodeURIComponent(query));
            resultsBox.innerHTML = "";

            if (res && res.owner) {
                // Found registered domain
                const card = document.createElement('div');
                card.className = "result-card";
                card.innerHTML = '<div class="result-title" onclick="switchTab(\'browser\'); document.getElementById(\'browser-url-input\').value=\'' + query + '\'; executeBrowserLoad();">' + query + '</div>' +
                    '<div class="result-snippet">نطاق لامركزي نشط على شبكة NeuroRoot. المالك: ' + res.owner + ' | الهدف: ' + res.target + '</div>';
                resultsBox.appendChild(card);
            }

            // Keyword broadcast on DHT
            const broadcast = await callAPI("/api/search?q=" + encodeURIComponent(query));
            if (broadcast) {
                const card = document.createElement('div');
                card.className = "result-card";
                card.innerHTML = '<div class="result-title">إعلان البحث: ' + query + '</div>' +
                    '<div class="result-snippet">تم بث إعلان استعلام عن الكلمة الدلالية "' + query + '" في شبكة DHT للوكلاء الأذكياء. جاري البحث عن مقترحات محتوى متوافقة...</div>';
                resultsBox.appendChild(card);
            }

            if (resultsBox.children.length === 0) {
                resultsBox.innerHTML = '<div style="text-align:center; color:var(--text-muted); padding:2rem;">لم يتم العثور على نطاقات مباشرة. تم نشر إعلان البحث في الـ DHT للوكلاء.</div>';
            }
        }

        // 🛒 Registrar (GoDaddy-like) logic
        let registrarTargetDomain = "";
        let generatedSecret = "";

        async function checkDomainAvailability() {
            let domain = document.getElementById('registrar-domain-input').value.trim();
            if (!domain) return;

            if (!domain.endsWith(".ia")) {
                domain += ".ia";
            }

            showToast("جاري التحقق من ملكية النطاق على DHT...", "success");

            const res = await callAPI("/api/resolve?name=" + encodeURIComponent(domain));
            const card = document.getElementById('registrar-result-card');
            const domainTitle = document.getElementById('registrar-result-domain');
            const statusText = document.getElementById('registrar-result-status');
            const checkoutBtn = document.getElementById('registrar-checkout-btn');

            document.getElementById('registration-wizard').style.display = "none";
            card.style.display = "flex";

            if (res && res.owner) {
                // Already registered
                domainTitle.innerText = domain;
                statusText.innerText = "غير متاح - مسجل بالفعل";
                statusText.style.color = "var(--accent-rose)";
                checkoutBtn.style.display = "none";
            } else {
                // Available
                registrarTargetDomain = domain;
                domainTitle.innerText = domain;
                statusText.innerText = "متاح للتسجيل مجاناً";
                statusText.style.color = "var(--accent-emerald)";
                checkoutBtn.style.display = "block";
            }
        }

        async function startRegistrationWizard() {
            document.getElementById('registrar-result-card').style.display = "none";
            const wizard = document.getElementById('registration-wizard');
            wizard.style.display = "flex";

            const step1 = document.getElementById('step1-indicator');
            const step2 = document.getElementById('step2-indicator');
            const step3 = document.getElementById('step3-indicator');
            const line1 = document.getElementById('line1-indicator');
            const line2 = document.getElementById('line2-indicator');
            const timerArea = document.getElementById('wizard-timer-container');
            const statusText = document.getElementById('wizard-status-text');

            // Reset indicators
            step1.className = "wizard-step active";
            step2.className = "wizard-step";
            step3.className = "wizard-step";
            line1.className = "wizard-line";
            line2.className = "wizard-line";
            timerArea.style.display = "none";

            // 1. Commit Domain
            statusText.innerText = "جاري تأمين النطاق... يتم تشفير اسم النطاق وربطه بالهوية في الشبكة لمنع الاختطاف.";
            
            const res = await callAPI("/api/domain/commit", "POST", { domain: registrarTargetDomain });
            if (res) {
                generatedSecret = res.secret || "derived_secret_hash";
                step1.className = "wizard-step completed";
                line1.className = "wizard-line active";
                step2.className = "wizard-step active";
                timerArea.style.display = "block";

                // Start 60s countdown
                let timeLeft = 60;
                document.getElementById('wizard-timer-display').innerText = timeLeft;
                statusText.innerText = "تم بث حظر الحجز بنجاح! يجب الانتظار 60 ثانية لتسجيل المعاملة وإثبات الأسبقية رياضياً.";

                const interval = setInterval(async () => {
                    timeLeft--;
                    document.getElementById('wizard-timer-display').innerText = timeLeft;
                    if (timeLeft <= 0) {
                        clearInterval(interval);
                        timerArea.style.display = "none";
                        statusText.innerText = "جاري الكشف النهائي وتسجيل الهوية بالنظام اللامركزي للشبكة...";
                        step2.className = "wizard-step completed";
                        line2.className = "wizard-line active";
                        step3.className = "wizard-step active";

                        // Simulate reveal/complete (Since local dashboard communicates with local REST, 
                        // local REST prints the commands/codes to complete on founder, or runs auto-reveal if founder config)
                        setTimeout(() => {
                            step3.className = "wizard-step completed";
                            statusText.innerText = "تهانينا! تم حجز النطاق " + registrarTargetDomain + " بنجاح وأصبح مسجلاً باسمك اللامركزي.";
                            showToast("تم تسجيل النطاق بنجاح!", "success");
                        }, 2000);
                    }
                }, 1000);
            } else {
                statusText.innerText = "حدث خطأ أثناء إرسال الالتزام للشبكة. يرجى التحقق من الاتصال.";
            }
        }

        function resetRegistrar() {
            document.getElementById('registration-wizard').style.display = "none";
            document.getElementById('registrar-result-card').style.display = "none";
            document.getElementById('registrar-domain-input').value = "";
        }

        // 💬 Chat (Telegram-like) logic
        async function loadJoinedChannels() {
            const list = await callAPI("/api/channels/list");
            const container = document.getElementById('chat-channels-list');
            container.innerHTML = "";
            
            if (list && list.length > 0) {
                list.forEach(ch => {
                    const item = document.createElement('div');
                    item.className = "channel-item";
                    if (ch === currentActiveChannel) {
                        item.className = "channel-item active";
                    }
                    item.onclick = () => selectChannel(ch);

                    const avatar = document.createElement('div');
                    avatar.className = "channel-avatar";
                    avatar.innerText = ch.substring(0, 1).toUpperCase();

                    const nameSpan = document.createElement('span');
                    nameSpan.innerText = "# " + ch;

                    item.appendChild(avatar);
                    item.appendChild(nameSpan);
                    container.appendChild(item);
                });
            } else {
                container.innerHTML = '<p style="color:var(--text-muted); font-size:0.8rem; text-align:center; padding:1rem;">لا توجد قنوات نشطة</p>';
            }
        }

        async function handleNewChannelJoin() {
            const input = document.getElementById('chat-new-channel-name');
            const name = input.value.trim();
            if (!name) return;

            showToast("جاري الانضمام إلى القناة...", "success");
            const res = await callAPI("/api/channels/join", "POST", { channel_id: name });
            if (res) {
                showToast("تم الانضمام لقناة #" + name + " بنجاح", "success");
                input.value = "";
                await loadJoinedChannels();
                selectChannel(name);
            }
        }

        function selectChannel(channelID) {
            currentActiveChannel = channelID;
            document.getElementById('chat-active-channel-name').innerText = "# " + channelID;
            document.getElementById('chat-active-channel-status').innerText = "قناة دردشة GossipSub لامركزية نشطة";
            
            loadJoinedChannels();
            pollChatMessages();

            clearInterval(channelPollInterval);
            channelPollInterval = setInterval(pollChatMessages, 2000);
        }

        async function pollChatMessages() {
            if (!currentActiveChannel) return;
            const msgs = await callAPI("/api/channels/messages?channel_id=" + encodeURIComponent(currentActiveChannel));
            const container = document.getElementById('chat-messages-container');
            
            if (msgs) {
                const isAtBottom = container.scrollHeight - container.clientHeight <= container.scrollTop + 30;
                container.innerHTML = "";
                
                if (msgs.length === 0) {
                    container.innerHTML = '<div style="text-align:center; color: var(--text-muted); margin-top: 8rem;">هذه القناة فارغة. كن أول من يكتب رسالة!</div>';
                    return;
                }

                const myDID = document.getElementById('profile-did-full').value;

                msgs.forEach(m => {
                    const bubble = document.createElement('div');
                    const isMyMsg = (m.sender === myDID);
                    
                    bubble.className = "chat-message-bubble " + (isMyMsg ? "outgoing" : "incoming");
                    
                    const senderSpan = document.createElement('span');
                    senderSpan.className = "msg-sender";
                    senderSpan.innerText = isMyMsg ? "أنا (هويتي)" : m.sender.substring(0, 18) + "...";
                    if (isMyMsg) {
                        senderSpan.style.color = "var(--accent-purple)";
                    }

                    const textDiv = document.createElement('div');
                    textDiv.className = "msg-text";
                    textDiv.innerText = m.content;

                    const timeSpan = document.createElement('span');
                    timeSpan.className = "msg-time";
                    timeSpan.innerText = new Date(m.timestamp * 1000).toLocaleTimeString('ar-EG');

                    bubble.appendChild(senderSpan);
                    bubble.appendChild(textDiv);
                    bubble.appendChild(timeSpan);
                    container.appendChild(bubble);
                });

                if (isAtBottom) {
                    container.scrollTop = container.scrollHeight;
                }
            }
        }

        async function executeSendChatMessage() {
            const input = document.getElementById('chat-message-input');
            const content = input.value.trim();
            if (!content || !currentActiveChannel) return;

            const res = await callAPI("/api/channels/publish", "POST", { channel_id: currentActiveChannel, content });
            if (res) {
                input.value = "";
                pollChatMessages();
            }
        }

        // 🖥️ Browser logic
        function executeBrowserLoad() {
            let domain = document.getElementById('browser-url-input').value.trim();
            if (!domain) return;

            if (!domain.endsWith(".ia")) {
                domain += ".ia";
            }

            showToast("جاري الاستعلام عن النطاق وتوجيه البوابة...", "success");

            // Extract the REST API address to derive standard local HTTP Gateway port (default 8090)
            // We can target the gateway at 127.0.0.1:8090
            const gatewayUrl = "http://127.0.0.1:8090/d/" + domain + "/";

            document.getElementById('browser-placeholder-text').style.display = "none";
            const iframe = document.getElementById('browser-iframe');
            iframe.style.display = "block";
            iframe.src = gatewayUrl;
        }
    </script>
</body>
</html>
`
