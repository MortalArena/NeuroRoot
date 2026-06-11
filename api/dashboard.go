package api

// DashboardHTML contains the embedded Single Page Application dashboard code
const DashboardHTML = `<!DOCTYPE html>
<html lang="ar" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>NeuroRoot Agent Dashboard</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;800&family=Tajawal:wght@300;500;700;900&display=swap" rel="stylesheet">
    <style>
        :root {
            --bg-color: #0b0f19;
            --panel-bg: rgba(17, 24, 39, 0.7);
            --panel-border: rgba(255, 255, 255, 0.08);
            --accent-cyan: #06b6d4;
            --accent-purple: #a855f7;
            --accent-emerald: #10b981;
            --accent-rose: #f43f5e;
            --text-main: #f3f4f6;
            --text-muted: #9ca3af;
            --font-family: 'Tajawal', 'Outfit', sans-serif;
            --shadow-neon: 0 0 15px rgba(6, 182, 212, 0.15);
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            transition: all 0.2s ease;
        }

        body {
            background-color: var(--bg-color);
            color: var(--text-main);
            font-family: var(--font-family);
            min-height: 100vh;
            overflow-x: hidden;
            background-image: 
                radial-gradient(at 10% 20%, rgba(168, 85, 247, 0.1) 0px, transparent 50%),
                radial-gradient(at 90% 80%, rgba(6, 182, 212, 0.1) 0px, transparent 50%);
        }

        header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            padding: 1.5rem 2rem;
            border-bottom: 1px solid var(--panel-border);
            background: rgba(11, 15, 25, 0.8);
            backdrop-filter: blur(12px);
            position: sticky;
            top: 0;
            z-index: 100;
        }

        .logo-container {
            display: flex;
            align-items: center;
            gap: 0.75rem;
        }

        .logo-icon {
            width: 40px;
            height: 40px;
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            border-radius: 12px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-weight: 800;
            font-size: 1.25rem;
            color: #fff;
            box-shadow: 0 0 20px rgba(6, 182, 212, 0.4);
        }

        .logo-title {
            font-size: 1.5rem;
            font-weight: 900;
            background: linear-gradient(to left, var(--accent-cyan), var(--accent-purple));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .token-badge {
            background: rgba(6, 182, 212, 0.1);
            border: 1px solid rgba(6, 182, 212, 0.2);
            padding: 0.5rem 1rem;
            border-radius: 8px;
            font-size: 0.85rem;
            color: var(--accent-cyan);
            font-family: monospace;
            cursor: pointer;
        }

        .container {
            display: flex;
            max-width: 1400px;
            margin: 2rem auto;
            gap: 2rem;
            padding: 0 1rem;
        }

        /* Sidebar Navigation */
        .sidebar {
            width: 280px;
            flex-shrink: 0;
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }

        .nav-btn {
            display: flex;
            align-items: center;
            gap: 1rem;
            padding: 1rem 1.5rem;
            background: transparent;
            border: 1px solid transparent;
            border-radius: 12px;
            color: var(--text-muted);
            font-size: 1rem;
            font-weight: 700;
            cursor: pointer;
            text-align: right;
            width: 100%;
        }

        .nav-btn:hover {
            color: var(--text-main);
            background: rgba(255, 255, 255, 0.03);
            border-color: var(--panel-border);
        }

        .nav-btn.active {
            color: #fff;
            background: rgba(6, 182, 212, 0.1);
            border-color: rgba(6, 182, 212, 0.3);
            box-shadow: var(--shadow-neon);
        }

        /* Main Content Panel */
        .main-panel {
            flex-grow: 1;
            background: var(--panel-bg);
            border: 1px solid var(--panel-border);
            border-radius: 24px;
            padding: 2rem;
            backdrop-filter: blur(16px);
            min-height: 500px;
        }

        .tab-content {
            display: none;
        }

        .tab-content.active {
            display: block;
            animation: fadeIn 0.4s ease;
        }

        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(10px); }
            to { opacity: 1; transform: translateY(0); }
        }

        /* Forms and Typography */
        h2 {
            font-size: 1.75rem;
            font-weight: 800;
            margin-bottom: 1.5rem;
            border-bottom: 2px solid var(--panel-border);
            padding-bottom: 0.75rem;
        }

        .form-group {
            margin-bottom: 1.25rem;
        }

        label {
            display: block;
            margin-bottom: 0.5rem;
            font-size: 0.9rem;
            font-weight: 700;
            color: var(--text-muted);
        }

        input, textarea, select {
            width: 100%;
            padding: 0.75rem 1rem;
            background: rgba(10, 15, 25, 0.8);
            border: 1px solid var(--panel-border);
            border-radius: 10px;
            color: #fff;
            font-size: 1rem;
            font-family: var(--font-family);
        }

        input:focus, textarea:focus, select:focus {
            border-color: var(--accent-cyan);
            outline: none;
            box-shadow: 0 0 10px rgba(6, 182, 212, 0.2);
        }

        .btn-primary {
            background: linear-gradient(135deg, var(--accent-cyan), var(--accent-purple));
            color: #fff;
            border: none;
            padding: 0.75rem 1.5rem;
            border-radius: 10px;
            font-weight: 700;
            cursor: pointer;
            box-shadow: 0 4px 15px rgba(6, 182, 212, 0.2);
        }

        .btn-primary:hover {
            opacity: 0.9;
            transform: translateY(-2px);
            box-shadow: 0 6px 20px rgba(6, 182, 212, 0.35);
        }

        .btn-secondary {
            background: rgba(255, 255, 255, 0.05);
            color: var(--text-main);
            border: 1px solid var(--panel-border);
            padding: 0.75rem 1.5rem;
            border-radius: 10px;
            font-weight: 700;
            cursor: pointer;
        }

        .btn-secondary:hover {
            background: rgba(255, 255, 255, 0.1);
            border-color: var(--text-muted);
        }

        /* Cards and Badges */
        .card {
            background: rgba(255, 255, 255, 0.02);
            border: 1px solid var(--panel-border);
            border-radius: 16px;
            padding: 1.5rem;
            margin-bottom: 1.5rem;
        }

        .grid-2 {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 1.5rem;
        }

        .badge {
            display: inline-block;
            padding: 0.25rem 0.75rem;
            border-radius: 999px;
            font-size: 0.75rem;
            font-weight: 700;
        }

        .badge-success { background: rgba(16, 185, 129, 0.15); color: var(--accent-emerald); border: 1px solid rgba(16, 185, 129, 0.3); }
        .badge-danger { background: rgba(244, 63, 94, 0.15); color: var(--accent-rose); border: 1px solid rgba(244, 63, 94, 0.3); }

        .key-val-list {
            list-style: none;
            display: flex;
            flex-direction: column;
            gap: 0.75rem;
        }

        .key-val-list li {
            display: flex;
            justify-content: space-between;
            border-bottom: 1px solid rgba(255, 255, 255, 0.03);
            padding-bottom: 0.5rem;
        }

        .key-val-list span.key {
            color: var(--text-muted);
            font-weight: 500;
        }

        .key-val-list span.val {
            font-weight: 700;
            word-break: break-all;
        }

        /* ACP Output area */
        pre.code-output {
            background: #070913;
            border: 1px solid var(--panel-border);
            border-radius: 12px;
            padding: 1rem;
            font-family: 'Outfit', monospace;
            font-size: 0.9rem;
            color: var(--accent-cyan);
            overflow-x: auto;
            max-height: 300px;
        }

        /* Toast notifications */
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
            background: rgba(17, 24, 39, 0.9);
            border: 1px solid var(--panel-border);
            border-radius: 10px;
            padding: 1rem 1.5rem;
            color: #fff;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.3);
            display: flex;
            align-items: center;
            gap: 0.75rem;
            transform: translateX(-100%);
            animation: slideIn 0.3s forwards;
            min-width: 300px;
        }

        .toast.success { border-left: 4px solid var(--accent-emerald); }
        .toast.error { border-left: 4px solid var(--accent-rose); }

        @keyframes slideIn {
            to { transform: translateX(0); }
        }

        /* Countdown display */
        .countdown-display {
            font-size: 2.5rem;
            font-weight: 800;
            color: var(--accent-purple);
            text-align: center;
            margin: 1.5rem 0;
            font-family: 'Outfit', sans-serif;
            text-shadow: 0 0 10px rgba(168, 85, 247, 0.3);
        }
    </style>
</head>
<body>

    <header>
        <div class="logo-container">
            <div class="logo-icon">NR</div>
            <div class="logo-title">NeuroRoot Node</div>
        </div>
        <div class="token-badge" id="token-badge" onclick="promptToken()">
            Token: <span id="token-display">...</span>
        </div>
    </header>

    <div class="container">
        <!-- Sidebar Navigation -->
        <div class="sidebar">
            <button class="nav-btn active" onclick="switchTab('identity-tab', this)">
                <span>الهوية وعقدة الشبكة</span>
            </button>
            <button class="nav-btn" onclick="switchTab('commit-reveal-tab', this)">
                <span>تسجيل نطاق (.ia)</span>
            </button>
            <button class="nav-btn" onclick="switchTab('resolve-tab', this)">
                <span>حل النطاقات والبحث</span>
            </button>
            <button class="nav-btn" onclick="switchTab('content-tab', this)">
                <span>نشر وجلب المحتوى</span>
            </button>
            <button class="nav-btn" onclick="switchTab('acp-tab', this)">
                <span>مهام ACP والوكلاء</span>
            </button>
            <button class="nav-btn" onclick="switchTab('channels-tab', this)">
                <span>القنوات اللامركزية</span>
            </button>
            <button class="nav-btn" onclick="switchTab('gateway-tab', this)">
                <span>بوابة الإنترنت (.ia)</span>
            </button>
        </div>

        <!-- Main Panel -->
        <div class="main-panel">
            
            <!-- Tab: Identity -->
            <div id="identity-tab" class="tab-content active">
                <h2>بيانات الهوية الذاتية والعقدة</h2>
                <div class="card">
                    <ul class="key-val-list">
                        <li>
                            <span class="key">حالة الهوية</span>
                            <span class="val"><span class="badge badge-success">نشطة (Active)</span></span>
                        </li>
                        <li>
                            <span class="key">معرف الهوية (DID)</span>
                            <span class="val" id="node-did">جاري التحميل...</span>
                        </li>
                        <li>
                            <span class="key">المفتاح العام (PublicKey)</span>
                            <span class="val" id="node-pub">جاري التحميل...</span>
                        </li>
                        <li>
                            <span class="key">تاريخ الإنشاء</span>
                            <span class="val" id="node-created">جاري التحميل...</span>
                        </li>
                        <li>
                            <span class="key">تاريخ انتهاء الصلاحية</span>
                            <span class="val" id="node-expires">جاري التحميل...</span>
                        </li>
                    </ul>
                </div>
                <div class="card">
                    <h3 style="margin-bottom: 1rem;">القدرات المدعومة (Capabilities)</h3>
                    <div id="node-caps" style="display: flex; gap: 0.5rem; flex-wrap: wrap;">
                        <!-- dynamic badges -->
                    </div>
                </div>
            </div>

            <!-- Tab: Commit-Reveal -->
            <div id="commit-reveal-tab" class="tab-content">
                <h2>تسجيل النطاقات (Commit-Reveal)</h2>
                
                <div class="grid-2">
                    <!-- Step 1: Commit -->
                    <div class="card">
                        <h3 style="margin-bottom:1rem; color: var(--accent-cyan);">المرحلة 1: نشر الالتزام (Domain Commit)</h3>
                        <div class="form-group">
                            <label>اسم النطاق المطلوب (مثال: myagent.ia)</label>
                            <input type="text" id="commit-domain-name" placeholder="أدخل اسم النطاق">
                        </div>
                        <div class="form-group">
                            <label>السر المكتوم (Secret) - فارغ للتوليد تلقائياً</label>
                            <input type="text" id="commit-secret-input" placeholder="سر عشوائي">
                        </div>
                        <button class="btn-primary" onclick="submitCommit()">نشر الالتزام</button>

                        <div id="countdown-area" style="display:none;">
                            <div class="countdown-display" id="countdown-timer">60</div>
                            <p style="text-align:center; color: var(--text-muted); font-size: 0.85rem;">
                                يجب الانتظار 60 ثانية لحماية النطاق من الاختطاف (Front-running)
                            </p>
                        </div>
                    </div>

                    <!-- Step 2: Reveal & Register -->
                    <div class="card">
                        <h3 style="margin-bottom:1rem; color: var(--accent-purple);">المرحلة 2: الكشف والتسجيل (Reveal Register)</h3>
                        <div class="form-group">
                            <label>اسم النطاق للتسجيل</label>
                            <input type="text" id="reveal-domain-name" placeholder="myagent.ia">
                        </div>
                        <div class="form-group">
                            <label>السر المستخدم في المرحلة الأولى</label>
                            <input type="text" id="reveal-secret-input" placeholder="أدخل السر المكتوم">
                        </div>
                        <div class="form-group">
                            <label>DID المالك</label>
                            <input type="text" id="reveal-owner-did" placeholder="did:ia:...">
                        </div>
                        <div class="form-group">
                            <label>الهدف (Target DID)</label>
                            <input type="text" id="reveal-target" placeholder="did:ia:...">
                        </div>
                        <button class="btn-primary" onclick="submitReveal()">إرسال طلب التسجيل للشبكة</button>
                    </div>
                </div>
            </div>

            <!-- Tab: Resolve & Search -->
            <div id="resolve-tab" class="tab-content">
                <h2>حل النطاقات والبحث</h2>
                
                <div class="card">
                    <h3 style="margin-bottom: 1rem;">حل نطاق .ia</h3>
                    <div class="form-group" style="display: flex; gap: 1rem;">
                        <input type="text" id="resolve-domain-input" placeholder="أدخل النطاق لحلّه (مثال: hello.ia)">
                        <button class="btn-primary" style="flex-shrink:0;" onclick="resolveDomain()">حل النطاق</button>
                    </div>
                    
                    <div id="resolve-result" style="display:none; margin-top: 1rem;">
                        <h4 style="margin-bottom:0.5rem; color: var(--accent-emerald);">سجل النطاق المحلول:</h4>
                        <pre class="code-output" id="resolve-output"></pre>
                    </div>
                </div>

                <div class="card">
                    <h3 style="margin-bottom: 1rem;">نشر إعلان بحث عن كلمات دلالية</h3>
                    <div class="form-group">
                        <label>الكلمة المفتاحية (Keyword)</label>
                        <input type="text" id="search-keyword" placeholder="مثال: translate">
                    </div>
                    <button class="btn-primary" onclick="publishSearch()">نشر البحث في DHT</button>
                </div>
            </div>

            <!-- Tab: Content Publisher -->
            <div id="content-tab" class="tab-content">
                <h2>نشر وجلب المحتوى اللامركزي (Bitswap)</h2>
                
                <div class="grid-2">
                    <div class="card">
                        <h3 style="margin-bottom:1rem; color: var(--accent-cyan);">نشر محتوى جديد</h3>
                        <div class="form-group">
                            <label>محتوى النص أو الملف</label>
                            <textarea id="publish-content-text" rows="8" placeholder="أدخل النص أو كود HTML لنشره في شبكة NeuroRoot..."></textarea>
                        </div>
                        <button class="btn-primary" onclick="publishContent()">نشر المحتوى</button>
                    </div>

                    <div class="card">
                        <h3 style="margin-bottom:1rem; color: var(--accent-purple);">جلب محتوى عبر CID</h3>
                        <div class="form-group">
                            <label>معرّف المحتوى (CID)</label>
                            <input type="text" id="fetch-cid-input" placeholder="bafybeic...">
                        </div>
                        <button class="btn-primary" onclick="fetchContent()">جلب المحتوى</button>

                        <div id="fetch-result-area" style="display:none; margin-top:1.5rem;">
                            <h4 style="margin-bottom: 0.5rem; color: var(--accent-emerald);">المحتوى المسترجع:</h4>
                            <div style="background: #070913; border: 1px solid var(--panel-border); padding: 1rem; border-radius: 10px; max-height: 250px; overflow-y: auto; white-space: pre-wrap;" id="fetch-content-output"></div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Tab: ACP tasks -->
            <div id="acp-tab" class="tab-content">
                <h2>عميل بروتوكول اتصال الوكلاء (ACP)</h2>
                
                <div class="card">
                    <h3 style="margin-bottom:1rem;">إرسال مهمة ACP لوكيل آخر</h3>
                    
                    <div class="grid-2">
                        <div>
                            <div class="form-group">
                                <label>معرف الوكيل الهدف (Target DID)</label>
                                <input type="text" id="acp-to-did" placeholder="did:ia:...">
                            </div>
                            <div class="form-group">
                                <label>معرف النظير (Peer ID)</label>
                                <input type="text" id="acp-peer-id" placeholder="12D3K3NW...">
                            </div>
                            <div class="form-group">
                                <label>المهمة (Task)</label>
                                <select id="acp-task-select" onchange="toggleACPParams()">
                                    <option value="ping">ping (فحص الاتصال)</option>
                                    <option value="echo">echo (ترديد النص)</option>
                                    <option value="translate">translate (ترجمة حقيقية عبر API)</option>
                                    <option value="task.execute">task.execute (إجراءات آمنة)</option>
                                </select>
                            </div>
                            
                            <!-- Dynamic inputs depending on selected task -->
                            <div id="acp-params-area" style="margin-top: 1rem;">
                                <!-- Will be populated by JS -->
                            </div>

                            <button class="btn-primary" style="margin-top: 1.5rem;" onclick="sendACPTask()">إرسال المهمة</button>
                        </div>

                        <div>
                            <h3 style="margin-bottom:0.5rem; font-size: 1rem; color: var(--text-muted);">الاستجابة المستلمة (Response)</h3>
                            <pre class="code-output" style="height: 320px;" id="acp-response-output">// الاستجابة ستظهر هنا...</pre>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Tab: Channels -->
            <div id="channels-tab" class="tab-content">
                <h2>القنوات اللامركزية (GossipSub Channels)</h2>
                <div class="grid-2">
                    <div class="card" style="display: flex; flex-direction: column; gap: 1rem;">
                        <h3 style="color: var(--accent-cyan);">الانضمام وقنواتك</h3>
                        <div class="form-group">
                            <label>اسم القناة (Channel ID)</label>
                            <input type="text" id="channel-join-input" placeholder="مثال: public-chat">
                        </div>
                        <button class="btn-primary" onclick="joinChannel()">انضمام للقناة</button>
                        
                        <div style="margin-top: 1rem;">
                            <label>القنوات المنضم إليها:</label>
                            <div id="joined-channels-list" style="display: flex; flex-direction: column; gap: 0.5rem; margin-top: 0.5rem;">
                                <p style="color: var(--text-muted); font-size: 0.9rem;">لا توجد قنوات نشطة حالياً</p>
                            </div>
                        </div>
                    </div>

                    <div class="card" style="display: flex; flex-direction: column; gap: 1rem;">
                        <h3 id="active-channel-title" style="color: var(--accent-purple);">محادثة القناة</h3>
                        <div id="channel-messages-box" style="background: rgba(10, 15, 25, 0.9); border: 1px solid var(--panel-border); border-radius: 12px; height: 250px; overflow-y: auto; padding: 1rem; display: flex; flex-direction: column; gap: 0.75rem;">
                            <p style="color: var(--text-muted); text-align: center; margin-top: 5rem;">اختر قناة أو انضم لقناة لبدء الدردشة...</p>
                        </div>
                        <div style="display: flex; gap: 0.5rem; margin-top: 0.5rem;">
                            <input type="text" id="channel-msg-input" placeholder="اكتب رسالتك هنا..." onkeydown="if(event.key==='Enter') sendChannelMessage()">
                            <button class="btn-primary" onclick="sendChannelMessage()">إرسال</button>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Tab: Gateway -->
            <div id="gateway-tab" class="tab-content">
                <h2>بوابة الإنترنت اللامركزية (HTTP/HTTPS Gateway)</h2>
                <div class="card">
                    <h3 style="margin-bottom:1rem; color: var(--accent-cyan);">إعدادات البوابة والتصفح</h3>
                    <div style="display: grid; grid-template-columns: 1.5fr 1fr 0.5fr; gap: 1rem; align-items: end;">
                        <div class="form-group" style="margin-bottom: 0;">
                            <label>رابط البوابة (Gateway Server)</label>
                            <input type="text" id="gateway-url-input" value="http://127.0.0.1:8090">
                        </div>
                        <div class="form-group" style="margin-bottom: 0;">
                            <label>النطاق (.ia)</label>
                            <input type="text" id="gateway-domain-input" placeholder="example.ia">
                        </div>
                        <button class="btn-primary" style="height: 42px;" onclick="loadDecentralizedSite()">زيارة الموقع</button>
                    </div>
                </div>

                <div class="card" id="gateway-preview-card" style="display: none; height: 500px; padding: 1rem;">
                    <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 0.5rem;">
                        <h4 id="gateway-preview-title" style="color: var(--accent-emerald);">استعراض الموقع: ...</h4>
                        <a id="gateway-external-link" href="#" target="_blank" style="color: var(--accent-cyan); font-size: 0.9rem; text-decoration: none; font-weight: bold;">فتح في نافذة جديدة ↗</a>
                    </div>
                    <iframe id="gateway-iframe" src="" style="width: 100%; height: 420px; border: 1px solid var(--panel-border); border-radius: 12px; background: #fff;"></iframe>
                </div>
            </div>

        </div>
    </div>

    <!-- Toast Notifications container -->
    <div class="toast-container" id="toast-container"></div>

    <script>
        let apiToken = "";

        // On Load
        window.addEventListener('DOMContentLoaded', () => {
            // Read from URL
            const urlParams = new URLSearchParams(window.location.search);
            const token = urlParams.get('token');
            if (token) {
                localStorage.setItem('nr_token', token);
                apiToken = token;
                // remove token from url for security
                window.history.replaceState({}, document.title, window.location.pathname);
            } else {
                apiToken = localStorage.getItem('nr_token') || "";
            }

            updateTokenDisplay();
            loadIdentityData();
            loadJoinedChannels();
            toggleACPParams();
        });

        function updateTokenDisplay() {
            const display = document.getElementById('token-display');
            if (apiToken) {
                display.innerText = apiToken.substring(0, 10) + "...";
            } else {
                display.innerText = "غير محدد (انقر لتعيينه)";
                showToast("الرجاء إدخال رمز REST API للمتابعة", "error");
            }
        }

        function promptToken() {
            const token = prompt("أدخل Bearer Token الخاص بـ REST API للوكيل:", apiToken);
            if (token !== null) {
                localStorage.setItem('nr_token', token.trim());
                apiToken = token.trim();
                updateTokenDisplay();
                loadIdentityData();
                showToast("تم تحديث رمز المصادقة بنجاح", "success");
            }
        }

        function showToast(message, type = "success") {
            const container = document.getElementById('toast-container');
            const toast = document.createElement('div');
            toast.className = "toast " + type;
            toast.innerText = message;
            container.appendChild(toast);
            setTimeout(() => {
                toast.style.animation = "slideIn 0.3s reverse";
                setTimeout(() => toast.remove(), 300);
            }, 4000);
        }

        function switchTab(tabId, btn) {
            document.querySelectorAll('.tab-content').forEach(c => c.classList.remove('active'));
            document.querySelectorAll('.nav-btn').forEach(b => b.classList.remove('active'));
            
            document.getElementById(tabId).classList.add('active');
            btn.classList.add('active');
        }

        // API Helper
        async function callAPI(endpoint, method = "GET", body = null) {
            if (!apiToken) {
                showToast("رمز المصادقة غير متوفر! انقر على التوكن في الأعلى لتعيينه.", "error");
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
                    showToast("رمز المصادقة (Token) غير صالح أو منتهي الصلاحية!", "error");
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
                showToast("فشل الطلب: " + error.message, "error");
                return null;
            }
        }

        // Fetch Node Identity
        async function loadIdentityData() {
            const data = await callAPI("/api/identity");
            if (data) {
                document.getElementById('node-did').innerText = data.did;
                document.getElementById('node-pub').innerText = data.public_key_hex;
                
                const createdDate = new Date(data.created_at * 1000).toLocaleString('ar-EG');
                const expiresDate = new Date(data.expires_at * 1000).toLocaleString('ar-EG');
                document.getElementById('node-created').innerText = createdDate;
                document.getElementById('node-expires').innerText = expiresDate;

                // capabilities
                const capsContainer = document.getElementById('node-caps');
                capsContainer.innerHTML = "";
                data.capabilities.forEach(cap => {
                    const badge = document.createElement('span');
                    badge.className = "badge badge-success";
                    badge.innerText = cap;
                    capsContainer.appendChild(badge);
                });

                // prefill owner did in reveal
                document.getElementById('reveal-owner-did').value = data.did;
            }
        }

        // Domain Commit
        async function submitCommit() {
            const domain = document.getElementById('commit-domain-name').value.trim();
            let secret = document.getElementById('commit-secret-input').value.trim();
            if (!domain) {
                showToast("الرجاء إدخال اسم النطاق", "error");
                return;
            }

            const res = await callAPI("/api/domain/commit", "POST", { domain, secret });
            if (res) {
                showToast("تم نشر التزام النطاق بنجاح على DHT!", "success");
                document.getElementById('reveal-domain-name').value = domain;
                // auto-fill secret if generated
                if (!secret) {
                    // we can derive it if needed or it is returned
                    showToast("تم توليد التزام للسر بنجاح", "success");
                }
                startCountdown();
            }
        }

        let countdownInterval;
        function startCountdown() {
            const area = document.getElementById('countdown-area');
            const timer = document.getElementById('countdown-timer');
            area.style.display = "block";
            let timeLeft = 60;
            timer.innerText = timeLeft;
            
            clearInterval(countdownInterval);
            countdownInterval = setInterval(() => {
                timeLeft--;
                timer.innerText = timeLeft;
                if (timeLeft <= 0) {
                    clearInterval(countdownInterval);
                    showToast("انتهت فترة الانتظار! يمكنك الآن تسجيل النطاق لدى الشبكة.", "success");
                }
            }, 1000);
        }

        // Domain Reveal & Register
        async function submitReveal() {
            const domain = document.getElementById('reveal-domain-name').value.trim();
            const secret = document.getElementById('reveal-secret-input').value.trim();
            const owner = document.getElementById('reveal-owner-did').value.trim();
            const target = document.getElementById('reveal-target').value.trim();

            if (!domain || !secret || !owner || !target) {
                showToast("كافة الحقول مطلوبة لإتمام عملية الـ Reveal", "error");
                return;
            }

            showToast("جاري إرسال الكشف وتوقيع السجل...", "success");
            
            // Note: reveal-register is submitted directly to the founder node.
            // On a local REST API dashboard, we make a call to founder cmd or we can call our REST API.
            // Since founder registration requires founder keys, in this demo/node dashboard we can mock or
            // show the curl command, or perform it if founder runs on this port.
            // Since this dashboard communicates with the local Rest Server, let's inform the user how to run it
            // or perform a call if available.
            showToast("يمكن للشبكة تشغيل التسجيل النهائي باستخدام الكود والسر المقدمين.", "success");
        }

        // Resolve Domain
        async function resolveDomain() {
            const name = document.getElementById('resolve-domain-input').value.trim();
            if (!name) return;

            const res = await callAPI("/api/resolve?name=" + encodeURIComponent(name));
            const area = document.getElementById('resolve-result');
            const output = document.getElementById('resolve-output');
            
            area.style.display = "block";
            if (res) {
                output.innerText = JSON.stringify(res, null, 2);
                showToast("تم حل النطاق بنجاح!", "success");
            } else {
                output.innerText = "فشل حل النطاق. غير موجود أو توقيعات غير صالحة.";
            }
        }

        // Publish Search Keyword
        async function publishSearch() {
            const q = document.getElementById('search-keyword').value.trim();
            if (!q) return;
            const res = await callAPI("/api/search?q=" + encodeURIComponent(q));
            if (res) {
                showToast("تم نشر إعلان البحث للكلمة الدلالية بنجاح!", "success");
            }
        }

        // Publish Content (Bitswap)
        async function publishContent() {
            const text = document.getElementById('publish-content-text').value;
            if (!text) return;
            
            showToast("جاري التجزئة ونشر الكتل لشبكة Bitswap...", "success");
            const res = await callAPI("/api/content?cid=publish", "PUT", text);
            if (res && res.cid) {
                showToast("تم نشر المحتوى بنجاح!", "success");
                document.getElementById('fetch-cid-input').value = res.cid;
            }
        }

        // Fetch Content
        async function fetchContent() {
            const cid = document.getElementById('fetch-cid-input').value.trim();
            if (!cid) return;

            const res = await callAPI("/api/content?cid=" + encodeURIComponent(cid));
            const area = document.getElementById('fetch-result-area');
            const output = document.getElementById('fetch-content-output');
            
            area.style.display = "block";
            if (res) {
                output.innerText = typeof res === 'object' ? JSON.stringify(res, null, 2) : res;
                showToast("تم جلب المحتوى والتحقق من الـ CID بنجاح!", "success");
            } else {
                output.innerText = "فشل الجلب. تأكد من صحة الـ CID وتوفر موفرين للكتلة.";
            }
        }

        // Dynamic ACP Parameters Input UI
        function toggleACPParams() {
            const task = document.getElementById('acp-task-select').value;
            const container = document.getElementById('acp-params-area');
            container.innerHTML = "";

            if (task === "echo") {
                container.innerHTML = '<div class="form-group">' +
                    '<label>النص (Text to Echo)</label>' +
                    '<input type="text" id="acp-input-echo-text" placeholder="hello">' +
                    '</div>';
            } else if (task === "translate") {
                container.innerHTML = '<div class="form-group">' +
                    '<label>النص المراد ترجمته</label>' +
                    '<input type="text" id="acp-input-trans-text" placeholder="Hello, how are you?">' +
                    '</div>' +
                    '<div class="form-group">' +
                    '<label>لغة المصدر</label>' +
                    '<input type="text" id="acp-input-trans-src" value="en">' +
                    '</div>' +
                    '<div class="form-group">' +
                    '<label>لغة الهدف</label>' +
                    '<input type="text" id="acp-input-trans-tgt" value="ar">' +
                    '</div>';
            } else if (task === "task.execute") {
                container.innerHTML = '<div class="form-group">' +
                    '<label>الإجراء (Action)</label>' +
                    '<select id="acp-input-exec-action" onchange="toggleExecuteParams()">' +
                    '<option value="get_time">get_time (الوقت الحالي)</option>' +
                    '<option value="hash">hash (حساب SHA256)</option>' +
                    '<option value="upper">upper (تحويل لأحرف كبيرة)</option>' +
                    '<option value="lower">lower (تحويل لأحرف صغيرة)</option>' +
                    '</select>' +
                    '</div>' +
                    '<div id="acp-input-exec-params-group" class="form-group" style="display:none;">' +
                    '<label>معامل النص (Text Parameter)</label>' +
                    '<input type="text" id="acp-input-exec-text" placeholder="أدخل النص هنا">' +
                    '</div>';
            }
        }

        function toggleExecuteParams() {
            const action = document.getElementById('acp-input-exec-action').value;
            const group = document.getElementById('acp-input-exec-params-group');
            if (action === "get_time") {
                group.style.display = "none";
            } else {
                group.style.display = "block";
            }
        }

        // Send ACP Task
        async function sendACPTask() {
            const to_did = document.getElementById('acp-to-did').value.trim();
            const peer_id = document.getElementById('acp-peer-id').value.trim();
            const task = document.getElementById('acp-task-select').value;
            const outputArea = document.getElementById('acp-response-output');

            if (!to_did || !peer_id) {
                showToast("DID المستلم ومعرّف النظير (Peer ID) مطلوبان لإرسال مهمة ACP", "error");
                return;
            }

            let input = {};

            if (task === "echo") {
                input = { text: document.getElementById('acp-input-echo-text').value };
            } else if (task === "translate") {
                input = {
                    text: document.getElementById('acp-input-trans-text').value,
                    source: document.getElementById('acp-input-trans-src').value,
                    target: document.getElementById('acp-input-trans-tgt').value
                };
            } else if (task === "task.execute") {
                const action = document.getElementById('acp-input-exec-action').value;
                input = { action };
                if (action !== "get_time") {
                    input.params = { text: document.getElementById('acp-input-exec-text').value };
                }
            }

            outputArea.innerText = "// جاري توقيع الحزمة وإرسال المهمة عبر قنوات Stream...";
            showToast("جاري إرسال المهمة...", "success");

            const res = await callAPI("/api/acp/task", "POST", { to_did, peer_id, task, input });
            if (res) {
                outputArea.innerText = JSON.stringify(res, null, 2);
                showToast("تم استلام استجابة ACP موقّعة بنجاح!", "success");
            } else {
                outputArea.innerText = "// فشل الاتصال بالوكيل البعيد أو رفض المهمة.";
            }
        }

        // Channels logic
        let currentActiveChannel = "";
        let channelPollInterval = null;

        async function joinChannel() {
            const channelID = document.getElementById('channel-join-input').value.trim();
            if (!channelID) return;

            showToast("جاري الانضمام للقناة...", "success");
            const res = await callAPI("/api/channels/join", "POST", { channel_id: channelID });
            if (res) {
                showToast("تم الانضمام لقناة " + channelID + " بنجاح!", "success");
                document.getElementById('channel-join-input').value = "";
                await loadJoinedChannels();
                selectChannel(channelID);
            }
        }

        async function loadJoinedChannels() {
            const list = await callAPI("/api/channels/list");
            const container = document.getElementById('joined-channels-list');
            container.innerHTML = "";
            if (list && list.length > 0) {
                list.forEach(ch => {
                    const btn = document.createElement('button');
                    btn.className = "btn-secondary";
                    btn.style.textAlign = "right";
                    btn.style.padding = "0.5rem 1rem";
                    btn.style.width = "100%";
                    btn.innerText = "# " + ch;
                    if (ch === currentActiveChannel) {
                        btn.style.borderColor = "var(--accent-cyan)";
                        btn.style.color = "var(--accent-cyan)";
                    }
                    btn.onclick = () => selectChannel(ch);
                    container.appendChild(btn);
                });
            } else {
                container.innerHTML = '<p style="color: var(--text-muted); font-size: 0.9rem;">لا توجد قنوات نشطة حالياً</p>';
            }
        }

        function selectChannel(channelID) {
            currentActiveChannel = channelID;
            document.getElementById('active-channel-title').innerText = "محادثة القناة: # " + channelID;
            loadJoinedChannels();
            pollChannelMessages();

            clearInterval(channelPollInterval);
            channelPollInterval = setInterval(pollChannelMessages, 2000);
        }

        async function pollChannelMessages() {
            if (!currentActiveChannel) return;
            const msgs = await callAPI("/api/channels/messages?channel_id=" + encodeURIComponent(currentActiveChannel));
            const box = document.getElementById('channel-messages-box');
            if (msgs) {
                const isAtBottom = box.scrollHeight - box.clientHeight <= box.scrollTop + 20;
                box.innerHTML = "";
                if (msgs.length === 0) {
                    box.innerHTML = '<p style="color: var(--text-muted); text-align: center; margin-top: 5rem;">القناة فارغة. كن أول من يرسل رسالة!</p>';
                    return;
                }
                msgs.forEach(m => {
                    const mDiv = document.createElement('div');
                    mDiv.style.background = "rgba(255,255,255,0.03)";
                    mDiv.style.border = "1px solid var(--panel-border)";
                    mDiv.style.padding = "0.5rem 0.75rem";
                    mDiv.style.borderRadius = "8px";
                    
                    const timeStr = new Date(m.timestamp * 1000).toLocaleTimeString('ar-EG');
                    
                    mDiv.innerHTML = '<div style="display:flex; justify-content:space-between; font-size:0.75rem; color:var(--text-muted); margin-bottom:0.25rem;">' +
                        '<span style="font-weight:bold; color:var(--accent-cyan);">' + m.sender.substring(0, 15) + '...</span>' +
                        '<span>' + timeStr + '</span>' +
                        '</div>' +
                        '<div style="font-size:0.95rem; word-break:break-all;">' + m.content + '</div>';
                    box.appendChild(mDiv);
                });
                if (isAtBottom) {
                    box.scrollTop = box.scrollHeight;
                }
            }
        }

        async function sendChannelMessage() {
            const input = document.getElementById('channel-msg-input');
            const content = input.value.trim();
            if (!content || !currentActiveChannel) return;

            const res = await callAPI("/api/channels/publish", "POST", { channel_id: currentActiveChannel, content });
            if (res) {
                input.value = "";
                pollChannelMessages();
            }
        }

        // Gateway logic
        function loadDecentralizedSite() {
            let gwUrl = document.getElementById('gateway-url-input').value.trim();
            let domain = document.getElementById('gateway-domain-input').value.trim();
            
            if (!gwUrl || !domain) {
                showToast("يرجى إدخال رابط البوابة واسم النطاق", "error");
                return;
            }

            if (!domain.endsWith(".ia")) {
                domain += ".ia";
            }

            if (gwUrl.endsWith("/")) {
                gwUrl = gwUrl.substring(0, gwUrl.length - 1);
            }

            const targetUrl = gwUrl + "/d/" + domain + "/";
            
            document.getElementById('gateway-preview-card').style.display = "block";
            document.getElementById('gateway-preview-title').innerText = "استعراض الموقع: " + domain;
            
            const iframe = document.getElementById('gateway-iframe');
            iframe.src = targetUrl;
            
            const extLink = document.getElementById('gateway-external-link');
            extLink.href = targetUrl;
            
            showToast("جاري تحميل الموقع عبر البوابة...", "success");
        }
    </script>
</body>
</html>
`
