document.addEventListener("DOMContentLoaded", () => {
    const form = document.querySelector("form");
    const input = document.querySelector("input[name='query']");
    const resultsContainer = document.createElement("div");
    resultsContainer.classList.add("results-container");
    document.querySelector(".container").appendChild(resultsContainer);

    // Event listener for the search form
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const query = input.value.trim();
        if (!query) {
            showNotification("Please enter a search term.", "warning");
            return;
        }

        showLoading(true);
        resultsContainer.innerHTML = "";

        try {
            // Send POST request to Go backend
            const response = await fetch("/search", {
                method: "POST",
                headers: { "Content-Type": "application/x-www-form-urlencoded" },
                body: new URLSearchParams({ query })
            });

            if (!response.ok) {
                throw new Error(`Server responded with status ${response.status}`);
            }

            const html = await response.text();
            const parser = new DOMParser();
            const doc = parser.parseFromString(html, "text/html");
            const resultsList = doc.querySelector("ul");

            // Inject new results into dashboard
            if (resultsList) {
                resultsContainer.innerHTML = `
                    <h2>Results for "${query}"</h2>
                    ${resultsList.outerHTML}
                `;
                highlightTerms(resultsContainer, query);
            } else {
                resultsContainer.innerHTML = `<p>No results found.</p>`;
            }

            showNotification("Search completed successfully.", "success");
        } catch (err) {
            console.error("Search error:", err);
            showNotification("An error occurred during search. Check console for details.", "error");
        } finally {
            showLoading(false);
        }
    });

    // Simple UI loading overlay
    function showLoading(isLoading) {
        let overlay = document.querySelector(".loading-overlay");
        if (!overlay) {
            overlay = document.createElement("div");
            overlay.className = "loading-overlay";
            overlay.innerHTML = `<div class="spinner"></div>`;
            document.body.appendChild(overlay);
        }
        overlay.style.display = isLoading ? "flex" : "none";
    }

    // Highlight search term in results
    function highlightTerms(container, term) {
        const innerHTML = container.innerHTML.replace(
            new RegExp(`(${term})`, "gi"),
            `<span class="highlight">$1</span>`
        );
        container.innerHTML = innerHTML;
    }

    // Notification banner (toast-style)
    function showNotification(message, type = "info") {
        const existing = document.querySelector(".notification");
        if (existing) existing.remove();

        const note = document.createElement("div");
        note.className = `notification ${type}`;
        note.textContent = message;
        document.body.appendChild(note);

        setTimeout(() => note.classList.add("visible"), 50);
        setTimeout(() => {
            note.classList.remove("visible");
            setTimeout(() => note.remove(), 500);
        }, 3000);
    }
});
