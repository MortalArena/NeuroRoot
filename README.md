# NeuroRoot Core

نواة شبكة وكلاء لامركزية (Agent Internet) — Layer 0.

## الميزات

- **هويات ذاتية السيادة** (`did:ia:...`) مع Ed25519 و BIP39 Mnemonic
- **نظام أسماء `.ia`** مع توقيع مزدوج (مؤسس + مالك)
- **قنوات عامة** (GossipSub) و **خاصة** (AES-256-GCM)
- **مراسلة مباشرة** مشفرة (NaCl box) مع تجزئة الملفات
- **بحث موزع** مع توقيعات و Rate Limiting
- **محتوى قابل للعنونة** (CID + Bitswap)
- **أمان**: PoW (scrypt), Revocation, Delegation, DHT Validators, Replay protection

## البناء

```bash
make build
make test
```

## التشغيل

### عقدة بذرة (Bootstrap)

```bash
make run-seed
# أو
go run ./cmd/seed -port 4001
```

### وكيل (Agent)

```bash
go run ./cmd/agent -port 4002 -bootstrap "/ip4/127.0.0.1/tcp/4001/p2p/<PEER_ID>"
```

### مؤسس (تسجيل نطاقات)

```bash
go run ./cmd/founder -action register -domain example.ia -owner did:ia:... -founder-key <hex>
```

## متغيرات البيئة

| المتغير | الوصف | الافتراضي |
|---------|-------|-----------|
| `NR_LISTEN_PORT` | منفذ الاستماع | 4001 |
| `NR_DATA_DIR` | مجلد البيانات | ./data |
| `NR_POW_DIFFICULTY` | صعوبة PoW | 10 |
| `NR_REST_PORT` | منفذ REST API | 8080 |
| `NR_STORAGE_QUOTA_MB` | حصة التخزين | 1024 |
| `NR_FOUNDER_PUB` | مفتاح المؤسس العام (hex) | — |

## المرحلة 4 — ميزات جديدة (Phase 4)

### لوحة التحكم المحلية (Web Dashboard)
- واجهة مستخدم رسومية متطورة (HTML/CSS/JS) مدمجة في خادم REST API.
- تُعرض تلقائياً على المسار: `http://127.0.0.1:8080/dashboard`
- تدعم إدارة الهوية، ونظام Commit-Reveal لتسجيل النطاقات مع عداد تنازلي (60 ثانية)، ونشر وجلب كتل المحتوى، وإرسال مهام ACP للوكلاء الآخرين.

### دعم TLS/HTTPS للبوابة (HTTP Gateway)
- خيار تشغيل البوابة بشكل آمن عبر بروتوكول HTTPS.
- إمكانية التوليد التلقائي لشهادات TLS ذاتية التوقيع (ECDSA P-256) أو استخدام شهادات مخصصة.
- خيارات التشغيل:
  ```bash
  go run ./cmd/gateway -port 8443 -tls -bootstrap "/ip4/127.0.0.1/tcp/4001/p2p/<PEER>"
  ```

### ربط محرك ترجمة حقيقي لـ ACP
- ربط مهمة `translate` بـ MyMemory API الفعلي لتقديم خدمة ترجمة نصوص فورية بين الوكلاء بدلاً من المعالجة الصورية (stub).

### إصلاح التوافق الرياضي للتشفير (Ed25519 → Curve25519)
- تصحيح ثغرة رياضية حرجة في المراسلة المباشرة وتشفير القنوات عبر استخدام حزمة `edwards25519` لتحويل المفاتيح العامة (Twisted Edwards to Montgomery coords) وحساب SHA-512 مع Clamping للمفاتيح الخاصة لضمان نجاح ECDH الفعلي.

## المرحلة 3 — ميزات جديدة

### Key Rotation للقنوات الخاصة
- عند طرد عضو: `KeyVersion` يزداد ومفتاح AES جديد يُوزَّع على الأعضاء المتبقين
- `MemberKeys`: مفتاح مشفّر لكل عضو
- `pkg/channel/rotation.go`

### HTTP Gateway
```bash
go run ./cmd/gateway -port 8090 -bootstrap "/ip4/127.0.0.1/tcp/4001/p2p/<PEER>" -founder-pub <hex>
# زيارة: http://127.0.0.1:8090/d/example.ia/
```

### مهام ACP إضافية
- `translate` — ترجمة (stub جاهز للربط)
- `task.execute` — إجراءات آمنة: `get_time`, `hash`, `upper`, `lower`

---

## المرحلة 2 — ميزات جديدة

### Keystore مشفّر
```bash
go run ./cmd/agent -init -data ./agent-data
# يولّد mnemonic ويحفظ المفتاح مشفّراً (scrypt + AES-GCM) في identity.key
```

### Commit-Reveal لتسجيل النطاقات
```bash
# 1. المالك ينشر التزام (الاسم مخفي)
go run ./cmd/agent -commit-domain example.ia -data ./agent-data

# 2. بعد 60 ثانية — المؤسس يسجّل
go run ./cmd/founder -action reveal-register -domain example.ia -owner did:ia:... -secret <SECRET> -founder-key <hex>
```

### ACP v1 (Agent Communication Protocol)
```bash
# المهام المدمجة: ping, echo
curl -H "Authorization: Bearer <token>" http://127.0.0.1:8080/api/acp/tasks
curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" \
  -d '{"to_did":"did:ia:...","peer_id":"...","task":"echo","input":{"text":"hello"}}' \
  http://127.0.0.1:8080/api/acp/task
```

## REST API

يعمل على `127.0.0.1` فقط مع مصادقة Bearer token:

- `GET /api/identity` — سجل الهوية
- `GET /api/search?q=` — نشر إعلان بحث
- `GET /api/resolve?name=` — حل نطاق `.ia`
- `GET /api/content?cid=` — جلب محتوى
- `PUT /api/content` — نشر محتوى
- `GET /api/acp/tasks` — المهام المدعومة
- `POST /api/acp/task` — إرسال مهمة ACP
- `POST /api/domain/commit` — نشر التزام نطاق
- `GET /api/domain/commit?commitment=` — جلب التزام

## Docker

```bash
docker compose -f docker/docker-compose.yml up
```

## الأمان

- جميع التوقيعات تستخدم **domain separation tags**
- المفاتيح الخاصة **لا تُنشر** على DHT
- REST API مقيد بـ localhost مع CORS صارم
- تحقق CID إلزامي عند جلب المحتوى

## الترخيص

MIT
